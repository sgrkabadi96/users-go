package routes

import (
	"github.com/gorilla/mux"
	"github.com/sgrkabadi96/bookAppClone/controllers"
)

func HandleUserRoutes(r *mux.Router) {

	r.HandleFunc("/user", controllers.GetAllUsers).Methods("GET")
	r.HandleFunc("/user/{id}", controllers.GetUserById).Methods("GET")
	r.HandleFunc("/user/{id}", controllers.DeleteUserById).Methods("DELETE")
	r.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", controllers.UpdateBookById).Methods("PUT")
	r.HandleFunc("/user", controllers.DeleteAllUser).Methods("DELETE")

}
