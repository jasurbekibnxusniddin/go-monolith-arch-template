package v1

import (
	"github.com/jasurbekibnxusniddin/go-monolith-arch-template/internal/pkgs/log"
	"github.com/jasurbekibnxusniddin/go-monolith-arch-template/internal/service"
)

type handler struct {
	service service.ServiceI
	log     log.Log
}

type Handler struct {
	Service service.ServiceI
	Log     log.Log
}

func NewHandler(service service.ServiceI, log log.Log) *handler {
	return &handler{service: service, log: log}
}
