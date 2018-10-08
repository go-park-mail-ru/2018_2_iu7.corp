package repositories

import (
	"2018_2_iu7.corp/service-registry/services/models"
)

type ServiceRepository interface {
	GetAllServicesInfo() (models.Services, error)
	GetServiceInfo(name string) (models.Service, error)

	RegisterService(name string, addr string) error
	UpdateService(name string, addr string) error
	UnregisterService(name string, addr string) error
}
