package models

type Task struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	UserId      int    `json:"userId"`
	Solved      bool   `json:"solved"`
}
