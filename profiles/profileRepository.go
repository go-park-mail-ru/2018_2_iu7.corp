package profiles

type ProfileRepository interface {
	SaveNew(p Profile) (err error)
	SaveExisting(p Profile) (err error)

	DeleteByID(id uint64) (err error)

	FindByID(id uint64) (p Profile, err error)
	FindByUsernameAndPassword(username, password string) (p Profile, err error)

	GetSeveralOrderByScorePaginated(page, pageSize int) (p []Profile, err error)
}
