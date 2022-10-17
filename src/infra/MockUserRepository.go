package infra

import "edwardhsu-golang-webapi/domain"

type MockUserRepository struct {
}

func NewMockUserRepository() domain.IUserRepository {
	return &MockUserRepository{}
}

func (r *MockUserRepository) GetUser(id string) *domain.User {
	return &domain.User{
		Name:   "User " + id,
		Gender: domain.Male,
	}
}
