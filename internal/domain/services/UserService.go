package services

import (
	"edwardhsu-golang-webapi/domain/models"
	"edwardhsu-golang-webapi/domain/repositories"
	"strconv"
	"time"
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

func (this *UserService) CreateUser(user *models.User) *models.User {
	user.UID = time.Now().Unix()
	user.EID = "A" + strconv.FormatInt(user.UID, 10)

	this.userRepository.SaveUser(user)

	return user
}

func (this *UserService) UpdateUser(user *models.User) *models.User {
	existsUser := this.FindUserByUid(user.UID)

	if existsUser == nil {
		return nil
	}

	existsUser.Status = user.Status

	this.userRepository.SaveUser(existsUser)

	return existsUser
}
