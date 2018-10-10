package rpc

import (
	"2018_2_iu7.corp/profile-service/models"
	"2018_2_iu7.corp/profile-service/repositories"
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/eureka"
)

type ProfileServiceImpl struct {
	profileRepository repositories.ProfileRepository
}

func CreateService(r repositories.ProfileRepository) (*micro.Service, error) {
	reg := eureka.NewRegistry(
		registry.Addrs("http://localhost:8761/eureka"),
	)

	service := micro.NewService(
		micro.Name(ServiceName),
		micro.Registry(reg),
	)
	service.Init()

	sImpl := &ProfileServiceImpl{
		profileRepository: r,
	}
	RegisterProfileServiceHandler(service.Server(), sImpl)

	return &service, nil
}

func (s *ProfileServiceImpl) GetProfile(ctx context.Context, req *Credentials, rsp *Profile) error {
	cr := models.Credentials{
		Username: req.Username,
		Password: req.Password,
	}

	p, err := s.profileRepository.FindByCredentials(cr)
	if err != nil {
		return err
	}

	rsp.ID = p.ProfileID
	return nil
}
