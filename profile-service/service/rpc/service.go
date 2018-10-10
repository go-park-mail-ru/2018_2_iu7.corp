package rpc

import (
	"2018_2_iu7.corp/profile-service/models"
	"2018_2_iu7.corp/profile-service/repositories"
	"context"
	"github.com/micro/go-micro"
)

type ProfileServiceImpl struct {
	profileRepository repositories.ProfileRepository
}

func CreateService(r repositories.ProfileRepository) *micro.Service {
	service := micro.NewService(
		micro.Name(ServiceName),
	)

	service.Init()
	RegisterProfileServiceHandler(service.Server(), new(ProfileServiceImpl))

	return &service
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
