package service

import (
	"errors"
	"log"
	"mlogreport/feature/admin/dto/request"
	"mlogreport/feature/admin/dto/response"
	"mlogreport/feature/admin/repository"
	"mlogreport/utils/auth"
	"mlogreport/utils/helper"
)

type adminService struct {
	adminRepository repository.AdminRepositoryInterface
}

type AdminServiceInterface interface {
	CreateAdvisor(data request.CreateAdvisor) error
	Login(data request.AdminLogin) (response.ResponseLogin, error)
}

func NewAdminService(adminRepository repository.AdminRepositoryInterface) AdminServiceInterface {
	return &adminService{
		adminRepository: adminRepository,
	}
}

func (admin *adminService) CreateAdvisor(data request.CreateAdvisor) error {
	err := admin.adminRepository.CreateAdvisor(data)
	if err != nil {
		return err
	}

	return nil
}

func (admin *adminService) Login(data request.AdminLogin) (response.ResponseLogin, error) {
	dataAdmin, err := admin.adminRepository.FindNip(data.Nip)
	if err != nil {
		return response.ResponseLogin{}, err
	}

	if !helper.CompareHash(dataAdmin.Password, data.Password) {
		return response.ResponseLogin{}, errors.New("error : password is wrong")
	}

	token, err := auth.CreateToken(dataAdmin.Nip, dataAdmin.Role)
	if err != nil {
		return response.ResponseLogin{}, err
	}

	log.Println("token",token)

	response := response.ModelToResponseLogin(dataAdmin.Name, dataAdmin.Role, token)
	return response, nil
}
