package routing

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"todo-mysql/handlers"
)

func Routing()  {
	r := mux.NewRouter()
	r.HandleFunc("/tasks", handlers.AddTasks).Methods("POST")
	r.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	r.HandleFunc("/task/{id}", handlers.GetTaskById).Methods("GET")
	r.HandleFunc("/task/{id}", handlers.DeleteTaskById).Methods("DELETE")
	r.HandleFunc("/task/{id}", handlers.UpdateTaskById).Methods("PUT")
	log.Fatal(http.ListenAndServe(":20001", r))
}
