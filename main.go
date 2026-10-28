package main

import (
	"fmt"
	"html/template"
	"net/http"
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
	// employees []domains.Employee

	template_html.ExecuteTemplate(writer, "Home.html", nil)
}
