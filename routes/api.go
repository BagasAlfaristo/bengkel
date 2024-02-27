package routes

import (
	"bengkel/app/http/controllers"
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	customers := controllers.GetAllCustomer()
	data := map[string]interface{}{
		"customers": customers,
	}

	temp, err := template.ParseFiles("resources/views/index.html")
	if err != nil {
		log.Fatal(err)
	}

	err = temp.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}
