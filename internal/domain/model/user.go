package model

import "time"

type UserData struct {
	Id          uint32    `json:"id"`
	Name        string    `json:"name"`
	Role        string    `json:"role"`
	Status      string    `json:"status"`
	Birthday    time.Time `json:"birthday"`
	PhoneNumber string    `json:"phoneNumber"`
}