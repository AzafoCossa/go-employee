package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// var template_htnl = template.Must(template.ParseGlob("views/*"))

func main(){
	fmt.Println("Server started at http://localhost:8000")
	http.HandleFunc("/", Home)

	http.ListenAndServe(":8000", nil)
}

func Home(writer http.ResponseWriter, request *http.Request){
	// template_htnl.ExecuteTemplate(writer, "Home", nil)
	html_template, error := template.ParseFiles("views/Home.html")

	if error != nil{
		fmt.Println("Could not parse the template")
	}

	html_template.Execute(writer, nil)
}
