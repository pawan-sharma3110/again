package main

import (
	"again/api/database"
	"again/api/handler"
	"net/http"
)

func main() {
	database.CreateUserTable(database.DB)
	http.HandleFunc("/register", handler.RegisterUser)
	http.HandleFunc("/user/delete/{userid}", handler.DeleteUser)
	http.HandleFunc("/user/login", handler.LoginUser)
	http.HandleFunc("/alluser", handler.GetAllUser)

	http.ListenAndServe(":8080", nil)
}
