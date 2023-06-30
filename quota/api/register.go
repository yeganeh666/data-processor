package api

import (
	models2 "IofIPOS/quota/api/src"
	"IofIPOS/quota/internal/configs"
	"IofIPOS/shared/grpcext"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	log           *log.Logger
	Config        *configs.Configs
	QuotaService  models2.UserQuotaServiceClient
	ObjectService models2.ObjectServiceClient
}

func NewService(log *log.Logger, conn *grpcext.Connection) *Service {

	return &Service{
		log:           log,
		QuotaService:  models2.NewUserQuotaServiceClient(conn),
		ObjectService: models2.NewObjectServiceClient(conn),
	}
}
