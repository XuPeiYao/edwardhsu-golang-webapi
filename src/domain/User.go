package domain

type Gender int

const (
	Male = iota
	Female
)

type User struct {
	Name   string `json:"name"`
	Gender Gender `json:"gender"`
}
