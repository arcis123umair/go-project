package models
type Todo struct {
	Id string `json:"id"`
	Tasks string `json:"tasks"`
	IsCompleted bool `json:"iscompleted"`
}

type Connect struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host string `json:"host"`
	Port int `json:"port"`
	Database string `json:"database"`
}