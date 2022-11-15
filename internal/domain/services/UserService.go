package services

import (
	"edwardhsu-golang-webapi/domain/models"
	"edwardhsu-golang-webapi/domain/repositories"
)

type UserService struct {
	userRepository repositories.IUserRepository
}

func NewUserService(userRepository repositories.IUserRepository) *UserService {
	service := &UserService{
		userRepository: userRepository,
	}

	return service
}

func (this *UserService) FindUserByUid(uid int64) *models.User {
	return this.userRepository.FindUserByUid(uid)
}
