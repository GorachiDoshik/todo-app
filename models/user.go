package models

import "errors"

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

func (i UpdateUser) Validate() error {
	if i.Name == nil && i.Username == nil && i.Password == nil && i.Email == nil {
		return errors.New("update structure has no value")
	}

	return nil
}
