package repositories

import (
	"2018_2_iu7.corp/profile-service/profiles/models"
)

type ProfileRepository interface {
	Open() (err error)
	Close() (err error)

	SaveNew(p models.Profile) (err error)
	SaveExisting(id uint32, u models.ProfileDataUpdate) (err error)

	DeleteByID(id uint32) (err error)

	FindByID(id uint32) (p models.Profile, err error)
	FindByUsernameAndPassword(username, password string) (p models.Profile, err error)

	GetSeveralOrderByScorePaginated(page, pageSize int) (p models.Profiles, err error)
}
