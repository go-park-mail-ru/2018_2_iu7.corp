package repositories

import (
	"2018_2_iu7.corp/profile-service/profiles/models"
)

type ProfileRepository interface {
	Open() (err error)
	Close() (err error)

	SaveNew(p models.Profile) (err error)
	SaveExisting(p models.Profile) (err error)

	DeleteByID(id int64) (err error)

	FindByID(id int64) (p models.Profile, err error)
	FindByUsernameAndPassword(username, password string) (p models.Profile, err error)

	GetSeveralOrderByScorePaginated(page, pageSize int32) (p models.Profiles, err error)
}
