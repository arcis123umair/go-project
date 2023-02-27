package models
type Todo struct {
	Id string `json:"id"`
	Tasks string `json:"tasks"`
	IsCompleted bool `json:"iscompleted"`
}