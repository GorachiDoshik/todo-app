package repository

import (
	"crypto/sha1"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/zhashkevych/todo-app/models"
)

const salt = "osdkdki202ds"

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) Update(userId int, input models.UpdateUser) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Username != nil {
		setValues = append(setValues, fmt.Sprintf("username=$%d", argId))
		args = append(args, *input.Username)
		argId++
	}

	if input.Password != nil {
		setValues = append(setValues, fmt.Sprintf("password_hash=$%d", argId))
		args = append(args, generatePasswordHash(*input.Password))
		argId++
	}

	if input.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email=$%d", argId))
		args = append(args, *input.Email)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", usersTable, setQuery, argId)

	args = append(args, userId)

	_, err := r.db.Exec(query, args...)

	return err
}

func (r *UserPostgres) Get(userId int) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT id, name, username, email FROM %s WHERE id = $1", usersTable)

	if err := r.db.Get(&user, query, userId); err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserPostgres) Delete(userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", usersTable)

	_, err := r.db.Exec(query, userId)

	return err
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
