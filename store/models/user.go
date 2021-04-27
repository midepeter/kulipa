package models

type User struct {
	ID              int    `json:"id"`
	Email            string `json:"email"`
	AssociationName  string `json:"association_name"`
	AssociationAlias string `json:"assocation_alias"`
	AssocationType   string `json:"association_type"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	ConfirmPassword  string `json:"confirm_password"`
}
