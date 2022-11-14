package infra

import (
	domainModels "edwardhsu-golang-webapi/domain/models"
	domainRepo "edwardhsu-golang-webapi/domain/repositories"
)

type MockUserRepository struct {
}

func NewMockUserRepository() domainRepo.IUserRepository {
	return &MockUserRepository{}
}

func (r *MockUserRepository) GetUser(id string) *domainModels.User {
	return &domainModels.User{
		Name:   "User " + id,
		Gender: domainModels.Male,
	}
}
