package models

import "time"

type User struct {
	UserId 			string
	UserName 		string
	UserEmail 	string
	Password 		string
	DateCreated time.Time
	DateUpdated time.Time
}