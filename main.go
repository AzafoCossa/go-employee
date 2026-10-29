package main

import (
	"fmt"
	"net/http"
	"text/template"

	"azafocossa.com/employee/domains"
)

var template_html = template.Must(template.ParseGlob("views/*"))

func main(){
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Println("Server started at http://localhost:8000")
	http.HandleFunc("/", Home)
	http.HandleFunc("/employees/insert", Insert)
	http.HandleFunc("/employees/delete", DeleteHandler)

	http.ListenAndServe(":8000", nil)
}

func Home(writer http.ResponseWriter, request *http.Request){
	var employees []domains.Employee
	employees = domains.GetEmployees()

	template_html.ExecuteTemplate(writer, "Home.html", employees)
}

func Insert(writer http.ResponseWriter, request *http.Request){
	var employee domains.Employee

	employee.FirstName = request.FormValue("firstName")
	employee.LastName = request.FormValue("lastName")
	employee.Salary = parseInt(request.FormValue("salary")) * 100
	employee.Department = request.FormValue("department")

	domains.InsertEmployee(employee)
	
	http.Redirect(writer, request, "/", http.StatusFound);
}

func DeleteHandler(writer http.ResponseWriter, request *http.Request){
	var employeeId int
	employeeId = parseInt(request.FormValue("id"))

	employee := domains.GetEmployeeById(employeeId)
	domains.DeleteEmployee(employee)

	http.Redirect(writer, request, "/", http.StatusFound)
}

func parseInt(value string) (int){
	var convertedValue int

	fmt.Sscanf(value, "%d", &convertedValue)

	return convertedValue
}
