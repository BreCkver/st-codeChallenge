package models

type Account struct {
	UserName string `bson:"userName" json:"userName,omitempty"`
	Email    string `bson:"email" json:"email,omitempty"`
}
