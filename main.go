package main

import (
	"fmt"
	"net/http"
	"text/template"

	"azafocossa.com/employee/domains"
	"github.com/gorilla/mux"
)

var template_html = template.Must(template.ParseGlob("views/*"))

func main(){
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	router := mux.NewRouter()
	fmt.Println("Server started at http://localhost:8000")

	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/employees/insert", InsertHandler)
	router.HandleFunc("/employees/update", UpdateHandler)
	router.HandleFunc("/employees/edit/{id}", EditHandler)
	router.HandleFunc("/employees/delete", DeleteHandler)
	http.Handle("/", router)
	http.ListenAndServe(":8000", nil)
}

func HomeHandler(writer http.ResponseWriter, request *http.Request){
	var employees []domains.Employee
	employees = domains.Employee{}.GetEmployees()

	template_html.ExecuteTemplate(writer, "Home.html", employees)
}

func EditHandler(writer http.ResponseWriter, request *http.Request){
	vars := mux.Vars(request)
	employeeId := vars["id"]
	employee := domains.Employee{}.GetEmployeeById(parseInt(employeeId))

	template_html.ExecuteTemplate(writer, "EditEmployee.html", employee)
}

func InsertHandler(writer http.ResponseWriter, request *http.Request){
	var employee domains.Employee

	employee.FirstName = request.FormValue("firstName")
	employee.LastName = request.FormValue("lastName")
	employee.Salary = parseInt(request.FormValue("salary")) * 100
	employee.Department = request.FormValue("department")

	domains.Employee.InsertEmployee(employee)
	
	http.Redirect(writer, request, "/", http.StatusFound);
}

func UpdateHandler(writer http.ResponseWriter, request *http.Request){
	var employee domains.Employee

	employee.Id = parseInt(request.FormValue("employeeId"))
	employee.Department = request.FormValue("department")
	employee.FirstName = request.FormValue("firstName")
	employee.LastName = request.FormValue("lastName")
	employee.Salary = parseInt(request.FormValue("salary")) * 100

	domains.Employee.UpdateEmployee(employee)

	http.Redirect(writer, request, "/", http.StatusFound)
}

func DeleteHandler(writer http.ResponseWriter, request *http.Request){
	var employeeId int
	employeeId = parseInt(request.FormValue("id"))

	employee := domains.Employee{}.GetEmployeeById(employeeId)
	domains.Employee.DeleteEmployee(employee)

	http.Redirect(writer, request, "/", http.StatusFound)
}

func parseInt(value string) (int){
	var convertedValue int

	fmt.Sscanf(value, "%d", &convertedValue)

	return convertedValue
}
