package models

type Payment struct {
	FullName  string `json:"fullname"`
	Email     string `json:"email"`
	Number    int    `json:"number"`
	MatricNo  int    `json:"matric_no"`
	Admission string `json:"admission"`
	Level     string `json:"level"`
}
