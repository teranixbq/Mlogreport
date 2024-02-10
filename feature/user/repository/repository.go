package repository

import (
	"errors"
	"mlogreport/feature/user/dto/request"
	"mlogreport/feature/user/dto/response"
	"mlogreport/feature/user/model"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepositoryInterface interface {
	InsertUser(data request.RequestUser) error
	FindNim(nim string) (model.Users, error)
	SelectUserById(nim string) (response.ProfileUser, error)
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &userRepository{
		db: db,
	}
}

var (
	pg *pgconn.PgError
)

func (user *userRepository) InsertUser(data request.RequestUser) error {
	input := request.RequestUserToModel(data)

	tx := user.db.Create(&input)

	if errors.As(tx.Error, &pg) {
		return errors.New("ERROR : data already exists")
	}

	if tx.Error != nil {
		return nil
	}

	return nil
}

func (user *userRepository) FindNim(nim string) (model.Users, error) {
	dataUser := model.Users{}

	tx := user.db.Where("nim = ?", nim).First(&dataUser)
	if tx.Error != nil {
		return model.Users{}, tx.Error
	}

	return dataUser, nil
}

func (user *userRepository) SelectUserById(nim string) (response.ProfileUser, error) {
	dataUser := model.Users{}

	tx := user.db.Where("nim = ?", nim).First(&dataUser)
	if tx.Error != nil {
		return response.ProfileUser{},tx.Error
	}

	response := response.ModelToProfileUser(dataUser)
	return response,nil
}
