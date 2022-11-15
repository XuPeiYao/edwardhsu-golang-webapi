package models

type User struct {
	UID     int64
	EID     string
	Status  Status
	Profile *UserProfile
}
