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
	http.HandleFunc("/employees/insert", Insert)

	http.ListenAndServe(":8000", nil)
}

func Home(writer http.ResponseWriter, request *http.Request){
	var employees []domains.Employee
	employees = domains.GetEmployees()

	log.Println(employees)

	template_html.ExecuteTemplate(writer, "Home.html", employees)
}

func Insert(writer http.ResponseWriter, request *http.Request){
	var employee domains.Employee

	employee.FirstName = request.FormValue("firstName")
	employee.LastName = request.FormValue("lastName")
	employee.Salary = parseInt(request.FormValue("salary")) * 100
	employee.Department = request.FormValue("department")

	domains.InsertEmployee(employee)

	employees := domains.GetEmployees()

	template_html.ExecuteTemplate(writer, "Home.html", employees)
}

func parseInt(value string) (int){
	var convertedValue int

	fmt.Sscanf(value, "%d", &convertedValue)

	return convertedValue
}
