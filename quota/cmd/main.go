package main

import (
	"IofIPOS/quota/internal/configs"
	"IofIPOS/quota/internal/repositories"
	"IofIPOS/quota/internal/services"
	"IofIPOS/shared/configext"
	"IofIPOS/shared/netext"
	"IofIPOS/shared/redisext"
	"IofIPOS/shared/service"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

func main() {
	conf := new(configs.Configs)
	loadedConfigs, err := configext.LoadConfigs("", configs.DefaultConfig, true, conf)
	if err != nil {
		log.Fatal(err)
	}
	conf = loadedConfigs.(*configs.Configs)
	logger := log.New()
	logger.SetReportCaller(true)

	redisClient := redisext.NewClient("", "")
	repository, err := repositories.NewRepository(logger, conf)
	if err != nil {
		panic("Failed to create repository")
	}

	quotaService := services.NewQuotaService(logger,
		redisClient,
		repositories.NewUserRepository(repository))

	objectService := services.NewObjectService(logger, quotaService,
		repositories.NewObjectRepository(repository))

	service.Serve(netext.Port(conf.GRPCPort), func(lst net.Listener) error {
		server := grpc.NewServer()
		quotaService.RegisterService(server)
		objectService.RegisterService(server)
		return server.Serve(lst)
	})

	service.Start(conf.ServiceName, conf.Version)

}
