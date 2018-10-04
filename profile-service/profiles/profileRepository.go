package profiles

type ProfileRepository interface {
	SaveNew(p Profile) (err error)
	SaveExisting(p Profile) (err error)

	DeleteByID(id int64) (err error)

	FindByID(id int64) (p Profile, err error)
	FindByUsernameAndPassword(username, password string) (p Profile, err error)

	GetSeveralOrderByScorePaginated(page, pageSize int32) (p []Profile, err error)
}
