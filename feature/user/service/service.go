package service

import (
	"errors"
	"mlogreport/feature/user/dto/request"
	"mlogreport/feature/user/dto/response"
	"mlogreport/feature/user/repository"
	"mlogreport/utils/auth"
	"mlogreport/utils/helper"
)

type userService struct {
	userRepository repository.UserRepositoryInterface
}

type UserServiceInterface interface {
	InsertUser(data request.RequestUser) error
	Login(data request.RequestLogin) (response.ResponseLogin, error)
}

func NewUserService(userRepository repository.UserRepositoryInterface) UserServiceInterface {
	return &userService{
		userRepository: userRepository,
	}
}

func (user *userService) InsertUser(data request.RequestUser) error {
	password, err := helper.HashPass(data.Password)
	if err != nil {
		return err
	}

	data.Password = password
	err = user.userRepository.InsertUser(data)
	if err != nil {
		return err
	}

	return nil
}

func (user *userService) Login(data request.RequestLogin) (response.ResponseLogin, error) {
	dataUser, err := user.userRepository.FindNim(data.Nim)
	if err != nil {
		return response.ResponseLogin{}, err
	}

	if !helper.CompareHash(dataUser.Password, data.Password) {
		return response.ResponseLogin{}, errors.New("error : password salah")
	}

	token, err := auth.CreateToken(dataUser.Nim, "")
	if err != nil {
		return response.ResponseLogin{}, err
	}

	response := response.ModelToResponseLogin(dataUser, token)
	return response, nil
}
