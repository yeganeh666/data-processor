package configext

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func LoadConfigs(path string, defaultConfig []byte, useDefault bool, config interface{}) (interface{}, error) {
	log.Infof("reding configs...")

	if useDefault {
		viper.SetConfigType("yaml")
		log.Infof("reading deafult configs")
		err := viper.ReadConfig(bytes.NewBuffer(defaultConfig))
		if err != nil {
			log.WithError(err).Error("read from default configs failed")
			return nil, err
		}
	} else {
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			log.Warnf("failed to read from env: %v", err)

			//viper.AddConfigPath("/etc/" + c.serviceName + "/")
			//viper.AddConfigPath("$HOME/." + c.serviceName)
			//viper.AddConfigPath(".")
			//viper.SetConfigName("config")

			viper.AddConfigPath("./configs")  //path for docker compose configs
			viper.AddConfigPath("../configs") //path for local configs
			viper.SetConfigName("config")
			viper.SetConfigType("yaml")
			if err = viper.ReadInConfig(); err != nil {
				log.Warnf("failed to read from yaml: %v", err)
				localErr := viper.ReadConfig(bytes.NewBuffer(defaultConfig))
				if localErr != nil {
					log.WithError(localErr).Error("read from default configs failed")
					return nil, localErr
				}
			}
		}
	}

	if err := viper.Unmarshal(config); err != nil {
		log.Errorf("faeiled unmarshal")
		return nil, err
	}

	return config, nil
}
