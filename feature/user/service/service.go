package service

import (
	"errors"

	"mlogreport/feature/user/dto/request"
	"mlogreport/feature/user/dto/response"
	"mlogreport/feature/user/repository"
	"mlogreport/utils/auth"
	"mlogreport/utils/constanta"
	"mlogreport/utils/helper"
	"mlogreport/utils/validation"
)

type userService struct {
	userRepository repository.UserRepositoryInterface
}

type UserServiceInterface interface {
	InsertUser(data request.RequestUser) error
	Login(data request.RequestLogin) (response.ResponseLogin, error)
	SelectUserById(id string) (response.ProfileUser, error)
	UpdateProfile(id string, data request.RequestUpdateProfile) error
}

func NewUserService(userRepository repository.UserRepositoryInterface) UserServiceInterface {
	return &userService{
		userRepository: userRepository,
	}
}

func (user *userService) InsertUser(data request.RequestUser) error {
	errEmpty := validation.CheckEmpty(data.Class, data.Name, data.Nim, data.Password)
	if errEmpty != nil {
		return errEmpty
	}

	errLength := validation.CheckLength(data.Password)
	if errLength != nil {
		return errLength
	}

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
	errEmpty := validation.CheckEmpty(data.Nim, data.Password)
	if errEmpty != nil {
		return response.ResponseLogin{}, errEmpty
	}

	dataUser, err := user.userRepository.FindNim(data.Nim)
	if err != nil {
		return response.ResponseLogin{}, err
	}

	if !helper.CompareHash(dataUser.Password, data.Password) {
		return response.ResponseLogin{}, errors.New(constanta.WRONG_PASS)
	}

	token, err := auth.CreateToken(dataUser.Id, "")
	if err != nil {
		return response.ResponseLogin{}, err
	}

	response := response.ModelToResponseLogin(dataUser, token)
	return response, nil
}

func (user *userService) SelectUserById(id string) (response.ProfileUser, error) {
	dataUser, err := user.userRepository.SelectUserById(id)
	if err != nil {
		return response.ProfileUser{}, err
	}
	return dataUser, nil
}

func (user *userService) UpdateProfile(id string, data request.RequestUpdateProfile) error {
	err := user.userRepository.UpdateProfile(id, data)
	if err != nil {
		return err
	}

	return nil
}
