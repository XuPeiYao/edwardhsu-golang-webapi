package models

import (
	domainModels "edwardhsu-golang-webapi/domain/models"
)

type UserProfile struct {
	Name             string  `json:"name"`
	Email            string  `json:"email"`
	ContactAddress   string  `json:"contact_address"`
	ContactTelephone string  `json:"contact_telephone"`
	ContactCellphone string  `json:"contact_cellphone"`
	AvatarUrl        *string `json:"avatar_url"`
}

func FromDomainUserProfile(user *domainModels.UserProfile) *UserProfile {
	return &UserProfile{
		Name:             user.Name,
		Email:            user.Email,
		ContactAddress:   user.ContactAddress,
		ContactTelephone: user.ContactTelephone,
		ContactCellphone: user.ContactCellphone,
		AvatarUrl:        user.AvatarUrl,
	}
}

func ToDomainUserProfile(user *UserProfile) *domainModels.UserProfile {
	return &domainModels.UserProfile{
		Name:             user.Name,
		Email:            user.Email,
		ContactAddress:   user.ContactAddress,
		ContactTelephone: user.ContactTelephone,
		ContactCellphone: user.ContactCellphone,
		AvatarUrl:        user.AvatarUrl,
	}
}
