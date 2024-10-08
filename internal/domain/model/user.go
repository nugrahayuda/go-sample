package model

type User struct {
	Id          uint32 `json:"id"`
	Name        string `json:"name"`
	RoleId      string `json:"role"`
	IsActive    string `json:"status"`
	Birthday    string `json:"birthday"`
	PhoneNumber string `json:"phoneNumber"`
}

func (User) TableName() string {
	return "users" // Set this to match your actual table name
}
