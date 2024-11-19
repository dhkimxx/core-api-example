package service

import (
	"core-api-example/dto"
	"core-api-example/entity"
	"core-api-example/repository"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func (us *UserService) CreateUser(userDTO dto.UserDTO) (*entity.User, error) {
	user := entity.User{
		Name:  userDTO.Name,
		Email: userDTO.Email,
	}
	return us.UserRepository.Save(user)
}
