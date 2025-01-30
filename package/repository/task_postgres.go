package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/zhashkevych/todo-app/models"
)

type TaskPostgres struct {
	db *sqlx.DB
}

func NewTaskPostgres(db *sqlx.DB) *TaskPostgres {
	return &TaskPostgres{db: db}
}

func (r *TaskPostgres) Create(userId int, task models.TaskCreateInput) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var taskId int
	createTaskQuery := fmt.Sprintf("INSERT INTO %s (user_id, title, description) VALUES ($1, $2, $3) RETURNING id", tasksTable)

	row := tx.QueryRow(createTaskQuery, userId, task.Title, task.Description)
	if err := row.Scan(&taskId); err != nil {
		return 0, nil
	}

	createTaskTagQuery := fmt.Sprintf("INSERT INTO %s (task_id, tag_id) VALUES ($1, $2)", taskTagTable)
	_, err = tx.Exec(createTaskTagQuery, taskId, task.Tagid)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createTaskCategoryQuery := fmt.Sprintf("INSERT INTO %s (task_id, category_id) VALUES ($1, $2)", taskCategoryTable)
	_, err = tx.Exec(createTaskCategoryQuery, taskId, task.Categoryid)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, nil
	}

	return taskId, nil
}

func (r *TaskPostgres) GetAll(userId int) ([]models.Task, error) {
	var tasks []models.Task
	query := fmt.Sprintf(`SELECT t.id, t.title, t.description, c.name as category, tg.name as tag FROM %s t
				INNER JOIN %s tc ON t.id = tc.task_id INNER JOIN %s c ON tc.category_id = c.id
				INNER JOIN %s tt ON t.id = tt.task_id INNER JOIN %s tg ON tt.tag_id = tg.id`,
		tasksTable, taskCategoryTable, categoriesTable, taskTagTable, tagsTable)

	err := r.db.Select(&tasks, query)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskPostgres) GetById(userId, taskId int) (models.Task, error) {
	var task models.Task
	query := fmt.Sprintf(`SELECT t.id, t.title, t.description, c.name as category, tg.name as tag FROM %s t
				INNER JOIN %s tc ON t.id = tc.task_id INNER JOIN %s c ON tc.category_id = c.id
				INNER JOIN %s tt ON t.id = tt.task_id INNER JOIN %s tg ON tt.tag_id = tg.id WHERE t.id = $1 AND t.user_id = $2`,
		tasksTable, taskCategoryTable, categoriesTable, taskTagTable, tagsTable)

	if err := r.db.Get(&task, query, taskId, userId); err != nil {
		return task, err
	}

	return task, nil
}

func (r *TaskPostgres) Update(userId, taskId int, input models.TaskUpdateInput) error {
	tx, err := r.db.Begin()
	if err != nil {
		return nil
	}

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d AND user_id=$%d", tasksTable, setQuery, argId, argId+1)

	args = append(args, taskId, userId)

	_, err = tx.Exec(query, args...)
	if err != nil {
		return err
	}

	if input.Categoryid != nil {
		query := fmt.Sprintf("UPDATE %s SET category_id=$1 WHERE task_id=$2", taskCategoryTable)
		_, err = tx.Exec(query, *input.Categoryid, taskId)
		if err != nil {
			return err
		}
	}

	if input.Tagid != nil {
		query := fmt.Sprintf("UPDATE %s SET tag_id=$1 WHERE task_id=$2", taskTagTable)
		_, err = tx.Exec(query, *input.Tagid, taskId)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *TaskPostgres) Delete(userId, taskId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 AND user_id = $2", tasksTable)
	_, err := r.db.Exec(query, taskId, userId)

	return err
}
