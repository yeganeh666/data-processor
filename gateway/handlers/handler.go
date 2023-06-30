package handlers

import (
	quota "IofIPOS/quota/api"
	"IofIPOS/shared/grpcext"
	"github.com/sirupsen/logrus"
)

type ObjectHandlerImpl struct {
	log     *logrus.Logger
	service *quota.Service
}

func NewObjectHandler(log *logrus.Logger, quotaAddr string) *ObjectHandlerImpl {
	quotaConn := grpcext.NewConnection(quotaAddr)
	return &ObjectHandlerImpl{
		log:     log,
		service: quota.NewService(log, quotaConn),
	}
}
