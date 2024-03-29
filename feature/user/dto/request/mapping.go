package request

import "mlogreport/feature/user/model"

func RequestUserToModel(data RequestUser) model.Users {
	return model.Users{
		Nim:      data.Nim,
		Name:     data.Name,
		Password: data.Password,
		Class:    data.Class,
	}
}

func ModelToUserUpdate(data RequestUpdateProfile) model.Users {
	return model.Users{
		Mitra:   data.Mitra,
		Program: data.Program,
	}
}
