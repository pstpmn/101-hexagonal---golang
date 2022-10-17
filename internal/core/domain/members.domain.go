package domain

import (
	"time"
)

type Members struct {
	Id         int       `json:"id"`
	FullName   string    `json:"fullName"`
	Username   string    `json:"username"`
	Password   string    `json:"-"`
	DateOfBird time.Time `json:"dateOfBird"`
	IsVip      bool      `json:"isVip"`
	CreatedAt  time.Time `json:"CreatedAt"`
}

func NewMember(id int, name string, user string, pass string, dob time.Time, isVip bool) *Members {
	return &Members{
		Id:         id,
		FullName:   name,
		Username:   user,
		Password:   pass,
		DateOfBird: dob,
		IsVip:      isVip,
	}
}
