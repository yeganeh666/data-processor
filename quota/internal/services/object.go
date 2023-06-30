package services

import (
	api "IofIPOS/quota/api/src"
	"IofIPOS/quota/internal/repositories"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type ObjectService interface {
	RegisterService(server *grpc.Server)
	PreUpload(ctx context.Context, req *api.PreUploadObjectReq) (*api.PreUploadObjectRes, error)
}

type ObjectServiceImpl struct {
	log              *log.Logger
	QuotaService     QuotaService
	ObjectRepository repositories.ObjectRepository
}

func NewObjectService(log *log.Logger, quotaService QuotaService,
	objectRepository repositories.ObjectRepository) ObjectService {
	return &ObjectServiceImpl{
		log:              log,
		ObjectRepository: objectRepository,
		QuotaService:     quotaService,
	}
}

func (s *ObjectServiceImpl) RegisterService(server *grpc.Server) {
	api.RegisterObjectServiceServer(server, s)
}

func (s *ObjectServiceImpl) PreUpload(ctx context.Context, req *api.PreUploadObjectReq) (*api.PreUploadObjectRes, error) {
	_, err := s.ObjectRepository.Get(req.Key, req.UserID)
	fmt.Println(err)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			s.log.WithError(err).Error("failed to do pre upload")
			return nil, err
		}

		_, err := s.QuotaService.CheckQuota(ctx, &api.CheckQuotaReq{DataSize: req.Size, UserID: req.UserID})
		if err != nil {
			s.log.WithError(err).Errorf("insufficient Quota for user %v", req.UserID)
			return nil, err
		}

		uploadKey := req.Key + uuid.New().String()
		userID, err := uuid.Parse(req.UserID)
		if err != nil {
			s.log.WithError(err).Error("failed to parse user ID")
			return nil, err
		}
		object := &repositories.Object{
			Key:       req.Key,
			UploadKey: uploadKey,
			UserID:    userID,
		}

		objectDetails, err := s.ObjectRepository.Create(object)
		if err != nil {
			s.log.WithError(err).Error("failed to create object details")
			return nil, err
		}

		return &api.PreUploadObjectRes{
			UploadID: objectDetails.UploadKey,
			UserID:   objectDetails.UserID.String(),
			Key:      objectDetails.Key,
		}, nil
	}

	return nil, errors.New("already exists")
}
