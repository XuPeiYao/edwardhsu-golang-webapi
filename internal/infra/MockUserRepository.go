package infra

import (
	domainModels "edwardhsu-golang-webapi/domain/models"
	domainRepo "edwardhsu-golang-webapi/domain/repositories"

	. "github.com/ahmetb/go-linq/v3"
)

type MockUserRepository struct {
	store []*domainModels.User
}

func NewMockUserRepository() domainRepo.IUserRepository {
	var instance = &MockUserRepository{
		store: make([]*domainModels.User, 0),
	}

	instance.store = append(instance.store,
		&domainModels.User{
			UID:    1,
			EID:    "A001",
			Status: domainModels.Enabled,
		},
		&domainModels.User{
			UID:    2,
			EID:    "A002",
			Status: domainModels.Disabled,
		},
		&domainModels.User{
			UID:    3,
			EID:    "A003",
			Status: domainModels.Enabled,
		},
	)

	return instance
}

func (this *MockUserRepository) FindUserByUid(uid int64) *domainModels.User {
	foundUser := From(this.store).FirstWithT(func(x *domainModels.User) bool {
		return x.UID == uid
	})

	if foundUser == nil {
		return nil
	}

	return foundUser.(*domainModels.User)
}
