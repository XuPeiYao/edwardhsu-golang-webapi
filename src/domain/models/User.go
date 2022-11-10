package models

type Gender int

const (
	Male   = 1
	Female = 2
)

type User struct {
	Name   string `json:"name"`
	Gender Gender `json:"gender"`
}
