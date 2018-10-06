package repositories

import (
	"2018_2_iu7.corp/profile-service/errors"
	"2018_2_iu7.corp/profile-service/profiles/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"strings"
)

const (
	DefaultHost     = "localhost"
	DefaultPort     = "5432"
	DefaultUser     = "postgres"
	DefaultPassword = ""
	DefaultDB       = "profiles"
)

type ConnectionParams struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type DBProfileRepository struct {
	db         *gorm.DB
	connParams ConnectionParams
}

func NewDBProfileRepository(connParams *ConnectionParams) *DBProfileRepository {
	return &DBProfileRepository{
		connParams: *connParams,
	}
}

func (r *DBProfileRepository) Open() (err error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		r.connParams.Host, r.connParams.Port, r.connParams.User, r.connParams.Password, r.connParams.Database)

	db, err := gorm.Open("postgres", connStr)
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
	m := &profileModel{}
	m.Profile = p

	errs := r.db.Create(m).GetErrors()
	if len(errs) != 0 {
		err := errs[0]
		if isConstraintViolationError(err) {
			return errors.NewConstraintViolationError("invalid profile: login or email is already taken")
		}
		return errors.NewServiceError()
	}

	return nil
}

func (r *DBProfileRepository) SaveExisting(p models.Profile) (err error) {
	return nil //TODO
}

func (r *DBProfileRepository) DeleteByID(id uint32) (err error) {
	return nil //TODO
}

func (r *DBProfileRepository) FindByID(id uint32) (p models.Profile, err error) {
	qModel := &profileModel{}
	qModel.ID = uint(id)

	var pModel profileModel
	if errs := r.db.Where(qModel).First(&pModel).GetErrors(); len(errs) != 0 {
		err := errs[0]
		if isNotFoundError(err) {
			return pModel.Profile, errors.NewNotFoundError("profile not found")
		}
		return pModel.Profile, errors.NewServiceError()
	}

	pModel.ProfileID = uint32(pModel.ID)
	return pModel.Profile, nil
}

func (r *DBProfileRepository) FindByUsernameAndPassword(username, password string) (p models.Profile, err error) {
	return models.Profile{}, nil //TODO
}

func (r *DBProfileRepository) GetSeveralOrderByScorePaginated(page, pageSize int) (p models.Profiles, err error) {
	return models.Profiles{}, nil //TODO
}

type profileModel struct {
	gorm.Model
	models.Profile
}

func isConstraintViolationError(err error) bool {
	return strings.Contains(err.Error(), "duplicate key value violates unique constraint")
}

func isNotFoundError(err error) bool {
	return strings.Contains(err.Error(), "record not found")
}
