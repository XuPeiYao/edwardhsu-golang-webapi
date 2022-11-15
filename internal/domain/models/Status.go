package models

// ref. https://www.sohamkamani.com/golang/enums/
type Status string

const (
	Disabled Status = "DISABLED"
	Enabled  Status = "ENABLED"
)
