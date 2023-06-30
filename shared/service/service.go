package service

import (
	"IofIPOS/shared/dispose"
	"IofIPOS/shared/netext"
	log "github.com/sirupsen/logrus"
)

func Start(name, version string) {
	var exposes []netext.Port
	interrupt := make(chan error)
	for port, serve := range serves {
		exposes = append(exposes, port)
		go func(port netext.Port, serve ServeFunc) {
			interrupt <- serve.Listen(port)
		}(port, serve)
	}

	data := log.Fields{
		"service": name,
		"version": version,
		"exposes": exposes,
	}
	log.WithFields(data).Info("service is running")

	interruptErr := <-interrupt
	if err := dispose.Close(); err != nil {
		log.WithError(err).Error("can not dispose")
	}
	if interruptErr == nil {
		log.WithFields(data).Panic("service interrupted")
	} else {
		log.WithFields(data).WithError(interruptErr).Fatal("service interrupted")
	}
}
