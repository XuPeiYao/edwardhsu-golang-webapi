package repositories

import "edwardhsu-golang-webapi/domain/models"

type IUserRepository interface {
	FindUserByUid(uid int64) *models.User
}
