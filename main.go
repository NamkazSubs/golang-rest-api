package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/NamkazSubs/golang-rest-api/config"
	"github.com/NamkazSubs/golang-rest-api/controllers"

	"github.com/gorilla/mux"
)

func main() {

	app := mux.NewRouter()
	app.Use(config.JwtAuthentication)
	//

	app.HandleFunc("/api/register", controllers.CreateAccount).Methods("POST")
	app.HandleFunc("/api/login", controllers.Authenticate).Methods("POST")
	app.HandleFunc("/", controllers.Home).Methods("GET")
	app.HandleFunc("/api/blog", controllers.AllPost).Methods("GET")
	app.HandleFunc("/api/blog/view", controllers.GetPost).Methods("POST")

	//Route Login
	app.HandleFunc("/api/admin/blog", controllers.AllPost).Methods("GET")
	app.HandleFunc("/api/admin/blog", controllers.CreatePost).Methods("POST")
	app.HandleFunc("/api/admin/blog", controllers.UpdatePost).Methods("PUT")
	app.HandleFunc("/api/admin/blog", controllers.DeletePost).Methods("DELETE")

	///
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, app)
	if err != nil {
		fmt.Print(err)
	}
}
