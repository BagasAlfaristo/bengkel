package routes

import (
	"bengkel/config"
	"log"
	"net/http"
)

func Web() {
	config.Dbconnection()
	http.HandleFunc("/", Index)

	log.Println("Server running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
