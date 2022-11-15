package models

type User struct {
	UID    int64  `json:"uid"`
	EID    string `json:"eid"`
	Status Status `json:"status"`
}
