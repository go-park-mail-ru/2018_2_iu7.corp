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
		err := classifyError(errs[0])
		return err
	}

	return nil
}

func (r *DBProfileRepository) SaveExisting(id uint32, u models.ProfileDataUpdate) (err error) {
	qModel := &profileModel{}
	qModel.ID = uint(id)

	if errs := r.db.Where(qModel).Update(u).GetErrors(); len(errs) != 0 {
		err := classifyError(errs[0])
		return err
	}

	return nil
}

func (r *DBProfileRepository) DeleteByID(id uint32) (err error) {
	return nil //TODO
}

func (r *DBProfileRepository) FindByID(id uint32) (p models.Profile, err error) {
	qModel := &profileModel{}
	qModel.ID = uint(id)

	var pModel profileModel
	if errs := r.db.Where(qModel).First(&pModel).GetErrors(); len(errs) != 0 {
		err := classifyError(errs[0])
		return pModel.Profile, err
	}

	pModel.ProfileID = uint32(pModel.ID)
	return pModel.Profile, nil
}

func (r *DBProfileRepository) FindByCredentials(cr models.Credentials) (p models.Profile, err error) {
	qModel := &profileModel{}
	qModel.Username = cr.Username
	qModel.Password = cr.Password

	var pModel profileModel
	if errs := r.db.Where(qModel).First(&pModel).GetErrors(); len(errs) != 0 {
		err := classifyError(errs[0])
		return pModel.Profile, err
	}

	pModel.ProfileID = uint32(pModel.ID)
	return pModel.Profile, nil
}

func (r *DBProfileRepository) GetSeveralOrderByScorePaginated(page, pageSize int) (p models.Profiles, err error) {
	return models.Profiles{}, nil //TODO
}

type profileModel struct {
	gorm.Model
	models.Profile
}

func classifyError(err error) error {
	msg := err.Error()

	if strings.Contains(msg, "duplicate key") {
		if strings.Contains(msg, "username") {
			return errors.NewConstraintViolationError("invalid username: already in use")
		}
		if strings.Contains(msg, "email") {
			return errors.NewConstraintViolationError("invalid username: already in use")
		}
	}

	if msg == "record not found" {
		return errors.NewNotFoundError("profile not found")
	}

	return errors.NewServiceError()
}
