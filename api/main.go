package main

import (
	"again/api/database"
	"again/api/handler"
	"again/api/middleware"
	"net/http"
)

func main() {
	database.CreateUserTable(database.DB)
	http.HandleFunc("/register", handler.RegisterUser)
	http.HandleFunc("/user/delete/{userid}", handler.DeleteUser)
	http.HandleFunc("/user/login", handler.LoginUser)
	http.Handle("/users", middleware.SecureApi(http.HandlerFunc(handler.GetAllUser)))
	http.ListenAndServe(":8080", nil)
}
