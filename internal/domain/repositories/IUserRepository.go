package repositories

import "edwardhsu-golang-webapi/domain/models"

type IUserRepository interface {
	GetUser(id string) *models.User
}
