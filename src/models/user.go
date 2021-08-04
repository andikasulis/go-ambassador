package models

type User struct {
	Id           uint
	FristName    string
	LastMame     string
	Email        string
	Password     []byte
	IsAmbassador bool
}
