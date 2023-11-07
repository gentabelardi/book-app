package main

import (
	"net/http"
    "github.com/gorilla/mux"
	"github.com/gentabelardi/book-app/controllers"
	"github.com/gentabelardi/book-app/models"
    "log"
    "fmt"
)

func main() {
    // Create a new gorilla mux router.
	models.ConnectDatabase()
    router := mux.NewRouter()

    // Register a simple route.
    router.HandleFunc("/books", controllers.GetAllBook).Methods("GET")
    router.HandleFunc("/books/{id}", controllers.GetBookByID).Methods("GET")
    router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
    router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
    router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")

    fmt.Println("server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
