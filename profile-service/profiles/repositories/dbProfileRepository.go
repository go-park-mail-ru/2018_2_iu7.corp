package repositories

import (
	"2018_2_iu7.corp/profile-service/errors"
	"2018_2_iu7.corp/profile-service/profiles/models"
	"github.com/jinzhu/gorm"
)

type profileModel struct {
	gorm.Model
	models.Profile
}

type DBProfileRepository struct {
	connStr string
	db      *gorm.DB
}

func NewDBProfileRepository(connStr string) *DBProfileRepository {
	return &DBProfileRepository{
		connStr: connStr,
	}
}

func (r *DBProfileRepository) Open() (err error) {
	db, err := gorm.Open("postgres", r.connStr)
	if err != nil {
		panic(err)
	}

	db = db.AutoMigrate(&profileModel{})
	if db == nil {
		return errors.NewServiceError()
	}

	r.db = db
	return nil
}

func (r *DBProfileRepository) Close() (err error) {
	return r.db.Close()
}

func (r *DBProfileRepository) SaveNew(p models.Profile) (err error) {
	return nil //TODO
}

func (r *DBProfileRepository) SaveExisting(p models.Profile) (err error) {
	return nil //TODO
}

func (r *DBProfileRepository) DeleteByID(id int64) (err error) {
	return nil //TODO
}

func (r *DBProfileRepository) FindByID(id int64) (p models.Profile, err error) {
	return models.Profile{}, nil //TODO
}

func (r *DBProfileRepository) FindByUsernameAndPassword(username, password string) (p models.Profile, err error) {
	return models.Profile{}, nil //TODO
}

func (r *DBProfileRepository) GetSeveralOrderByScorePaginated(page, pageSize int32) (p models.Profiles, err error) {
	return models.Profiles{}, nil //TODO
}
