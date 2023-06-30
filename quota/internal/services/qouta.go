package services

import (
	api "IofIPOS/quota/api/src"
	"IofIPOS/quota/internal/repositories"
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"strconv"
	"sync"
	"time"
)

type QuotaService interface {
	RegisterService(server *grpc.Server)
	CheckQuota(ctx context.Context, req *api.CheckQuotaReq) (*api.Void, error)
}
type QuotaServiceImpl struct {
	log            *log.Logger
	redis          *redis.Client
	redisLock      sync.Mutex
	UserRepository repositories.UserRepository
	rateLimiter    *rate.Limiter
}

func NewQuotaService(log *log.Logger,
	redisClient *redis.Client,
	userRepository repositories.UserRepository) QuotaService {
	return &QuotaServiceImpl{
		log:            log,
		redis:          redisClient,
		UserRepository: userRepository,
		rateLimiter:    rate.NewLimiter(rate.Limit(1), 1), //Limit of one request per second
	}
}

func (s *QuotaServiceImpl) RegisterService(server *grpc.Server) {
	api.RegisterUserQuotaServiceServer(server, s)
}

func (s *QuotaServiceImpl) CheckQuota(ctx context.Context, req *api.CheckQuotaReq) (*api.Void, error) {
	if err := s.rateLimiter.Wait(ctx); err != nil {
		s.log.WithError(err).Error("rate limit exceeded")
		return nil, err
	}
	userQuota, err := s.UserRepository.GetQuota(req.UserID)
	if err != nil {
		s.log.WithError(err).Error("user quota is not available")
		return nil, err
	}

	key := fmt.Sprintf("user:%v", req.UserID)

	s.redisLock.Lock()
	defer s.redisLock.Unlock()

	// Checking for the existence of a previous request and setting a valid value for last_request_time
	lastRequestTimeExists, err := s.redis.Exists(ctx, key+":last_request_time").Result()
	if err != nil {
		s.log.WithError(err).Error("failed to check last request time")
		return nil, err
	}

	currentTime := time.Now().Unix()
	var lastRequestTime int64

	if lastRequestTimeExists == 0 {
		// No previous request, set initial value for last_request_time
		_, err = s.redis.Set(ctx, key+":last_request_time", currentTime, 60*time.Second).Result()
		if err != nil {
			s.log.WithError(err).Error("failed to set last request time")
			return nil, err
		}
		lastRequestTime = currentTime
	} else {
		//Get the last_request_time value
		lastRequestTimeString, err := s.redis.Get(ctx, key+":last_request_time").Result()
		if err != nil {
			s.log.WithError(err).Error("failed to get last request time")
			return nil, err
		}
		lastRequestTime, err = strconv.ParseInt(lastRequestTimeString, 10, 64)
		if err != nil {
			s.log.WithError(err).Error("failed to parse last request time")
			return nil, err
		}
	}

	dataSizeMonthExists, err := s.redis.Exists(ctx, key+":data_size_month").Result()
	if err != nil {
		s.log.WithError(err).Error("failed to check last request time")
		return nil, err
	}
	if dataSizeMonthExists == 0 {
		// Setting the expiration time on the user volume key
		expiration := time.Until(time.Now().AddDate(0, 1, 0))
		err := s.redis.Expire(ctx, key+":data_size_month", expiration).Err()
		if err != nil {
			s.log.WithError(err).Errorf("error zeroing user %v consumption volume in month", userQuota.UserID)
			return nil, err
		}
	}

	if (currentTime - lastRequestTime) < int64(time.Minute/time.Second) {
		// The number of requests is in the time range, continue processing
		// Number of user requests per minute
		err = s.process(ctx, key, req.DataSize, userQuota)
		if err != nil {
			s.log.WithError(err).Errorf("failed to process request for %v user", req.UserID)
			return nil, err
		}
		userQuota.Lock()
		userQuota.CurrentDataVolume += req.DataSize
		userQuota.Unlock()
	} else {
		// The number of requests exceeds the limit, apply restrictions
		wg := sync.WaitGroup{}
		wg.Add(2)

		go func() {
			defer wg.Done()
			_, err := s.redis.Set(ctx, key+":last_request_time", currentTime, 60*time.Second).Result()
			if err != nil {
				s.log.WithError(err).Error("failed to set last request time")
			}
		}()

		go func() {
			defer wg.Done()
			_, err := s.redis.Set(ctx, key+":requests_minute", 1, 60*time.Second).Result()
			if err != nil {
				s.log.WithError(err).Errorf("error receiving %v user requests per minute", req.UserID)
			}
		}()

		wg.Wait()
	}

	return nil, s.UserRepository.SaveQuota(userQuota)
}

func (s *QuotaServiceImpl) process(ctx context.Context, key string, dataSize int64, userQuota *repositories.UserQuota) error {
	_ = s.redis.Expire(ctx, key+":requests_minute", 60*time.Second).Err()
	requestsPerMinute, err := s.redis.Incr(ctx, key+":requests_minute").Result()
	if err != nil {
		s.log.WithError(err).Errorf("error receiving %v user requests per minute: %s", userQuota.UserID, err.Error())
		return err
	}

	// Check the limitations
	if requestsPerMinute > int64(userQuota.RequestsPerMinute) {
		err = errors.New(fmt.Sprintf("user %v has exceeded the allowed number of requests per minute", userQuota.UserID))
		return err

	}

	// The total amount of data sent by the user per month
	dataSizePerMonth, err := s.redis.IncrBy(ctx, key+":data_size_month", dataSize).Result()
	if err != nil {
		s.log.WithError(err).Errorf("error in receiving the total amount of data sent by user %v in the month", userQuota.UserID)
		return err
	}

	if dataSizePerMonth >= userQuota.TotalDataVolumePerMonth {
		dataSizePerMonth, err = s.redis.DecrBy(ctx, key+":data_size_month", dataSize).Result()
		if err != nil {
			s.log.WithError(err).Errorf("error in receiving the total amount of data sent by user %v in the month", userQuota.UserID)
			return err
		}

		err = errors.New(fmt.Sprintf("user %v has exceeded the data limit", userQuota.UserID))
		return err
	}

	// Save the time of sending data
	err = s.redis.Set(ctx, key+":last_request_time", time.Now().Unix(), 0).Err()
	if err != nil {
		s.log.WithError(err).Errorf("error saving last request time for user %v", userQuota.UserID)
		return err
	}

	return nil
}
