package models

import "time"

type Todo struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at"`
}

type CreateTodoReq struct {
	Title string `json:"title"`
}

type UpdateTodoReq struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	CompletedAt time.Time `json:"completed_at"`
}

type GetListResp struct {
	Todos []*Todo
	Count int32
}
