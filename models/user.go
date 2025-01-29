package models

type User struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type UpdateUser struct {
	Name     *string `json:"name"`
	Username *string `json:"username"`
	Password *string `json:"password"`
	Email    *string `json:"email"`
}
