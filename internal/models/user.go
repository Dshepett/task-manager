package models

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Tasks []Task `json:"tasks"`
}