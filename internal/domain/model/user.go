package model

type UserData struct {
	Id          uint32 `json:"id"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	Status      string `json:"status"`
	Birthday    string `json:"birthday"`
	PhoneNumber string `json:"phoneNumber"`
}
