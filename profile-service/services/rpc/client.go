package rpc

import (
	"github.com/micro/go-micro"
)

func CreateClient() *ProfileService {
	service := micro.NewService(micro.Name(ClientName))
	service.Init()
	client := NewProfileService(ServiceName, service.Client())
	return &client
}
