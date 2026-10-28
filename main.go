package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"azafocossa.com/employee/domains"
)

var template_html = template.Must(template.ParseGlob("views/*"))

func main(){
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Println("Server started at http://localhost:8000")
	http.HandleFunc("/", Home)

	http.ListenAndServe(":8000", nil)
}

func Home(writer http.ResponseWriter, request *http.Request){
	var employees []domains.Employee
	employees = domains.GetEmployees()

	log.Println(employees)

	template_html.ExecuteTemplate(writer, "Home.html", employees)
}
