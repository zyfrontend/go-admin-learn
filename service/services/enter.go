package services

import (
	"app/services/back"
)

type ServiceGroup struct {
	BackServiceGroup back.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
