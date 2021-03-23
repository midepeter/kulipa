package models

type User struct {
	ID int `json:"id"`
	MatricNo int `json:"matric_no"`
	Email string `json:"email"`
}