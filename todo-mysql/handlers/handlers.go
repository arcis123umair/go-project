package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"todo-mysql/config"
	"todo-mysql/models"
)

func ConnectDatabase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var db models.Connect
	json.NewDecoder(r.Body).Decode(&db)
	config.Connect(db)
	json.NewEncoder(w).Encode("Connected to database")
}

func AddTasks(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var add []models.Todo
	json.NewDecoder(r.Body).Decode(&add)
	config.Database.Create(&add)
	json.NewEncoder(w).Encode(add)
}

func GetTasks(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var list []models.Todo
	config.Database.Find(&list)
	json.NewEncoder(w).Encode(list)
}

func GetTaskById(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var listid models.Todo
	id := mux.Vars(r)["id"]
	config.Database.First(&listid, id)
	json.NewEncoder(w).Encode(listid)
}

func DeleteTaskById(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var delete models.Todo
	id := mux.Vars(r)["id"]
	config.Database.Delete(&delete, id)
	json.NewEncoder(w).Encode(delete)
}

func UpdateTaskById(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var update models.Todo
	id := mux.Vars(r)["id"]
	config.Database.First(&update, id)
	fmt.Println(update)
	json.NewDecoder(r.Body).Decode(&update)
	err := config.Database.Save(&update)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(update)

}