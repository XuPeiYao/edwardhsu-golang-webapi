package models

import (
	domainModels "edwardhsu-golang-webapi/domain/models"
)

type User struct {
	UID     int64        `json:"uid"`
	EID     string       `json:"eid"`
	Status  string       `json:"status"`
	Profile *UserProfile `json:"profile"`
}

func FromDomainUser(user *domainModels.User) *User {
	domainInstance := &User{
		UID:    user.UID,
		EID:    user.EID,
		Status: string(user.Status),
	}

	if user.Profile != nil {
		domainInstance.Profile = FromDomainUserProfile(user.Profile)
	}

	return domainInstance
}

func ToDomainUser(user *User) *domainModels.User {
	appInstance := &domainModels.User{
		UID:    user.UID,
		EID:    user.EID,
		Status: domainModels.Status(user.Status),
	}

	if user.Profile != nil {
		appInstance.Profile = ToDomainUserProfile(user.Profile)
	}

	return appInstance
}
