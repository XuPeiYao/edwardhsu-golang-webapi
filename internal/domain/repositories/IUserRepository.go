package repositories

import domainModels "edwardhsu-golang-webapi/domain/models"

type IUserRepository interface {
	FindUserByUid(uid int64) *domainModels.User
	SaveUser(user *domainModels.User)
}
