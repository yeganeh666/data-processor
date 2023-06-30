package main

import (
	"IofIPOS/gateway/docs"
	"IofIPOS/gateway/handlers"
	"IofIPOS/gateway/internal/configs"
	"IofIPOS/gateway/router"
	"IofIPOS/shared/configext"
	"fmt"
	"github.com/sirupsen/logrus"
)

// @BasePath /api
// IofIPOS service godoc

func main() {
	log := logrus.New()

	conf := new(configs.Configs)
	loadedConfigs, err := configext.LoadConfigs("", configs.DefaultConfig, true, conf)
	if err != nil {
		log.Fatal(err)
	}
	conf = loadedConfigs.(*configs.Configs)

	initSwagger(conf)

	handler := handlers.NewObjectHandler(
		log, conf.QuotaAddress)

	router.Register(handler)
}

func initSwagger(conf *configs.Configs) {

	docs.SwaggerInfo.Title = "Input of Input-Processor-Output-Storage"
	docs.SwaggerInfo.Description = "IofIPOS Service : This is a simplified implementation of the data processor's input component."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("%v:%v", conf.HttpHost, conf.HttpPort)
	docs.SwaggerInfo.Schemes = []string{"http"}

}
