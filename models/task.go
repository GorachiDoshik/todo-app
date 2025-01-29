package models

type TaskCreateInput struct {
	Id          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Categoryid  int    `json:"category_id"`
	Tagid       int    `json:"tag_id"`
}

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Tag         string `json:"tag"`
}

type TaskUpdateInput struct {
	Title       *string `json:"title" db:"title"`
	Description *string `json:"description" db:"description"`
	Categoryid  *int    `json:"category_id"`
	Tagid       *int    `json:"tag_id"`
}
