package main

import (
	"fmt"
	"net/http"
	"text/template"

	"azafocossa.com/employee/domains"
	"github.com/gin-gonic/gin"
)

var template_html = template.Must(template.ParseGlob("views/*"))

func main(){
	router := gin.Default()
	router.Static("/static/", "./assets")
	router.LoadHTMLGlob("views/*")
	router.GET("/", HomeHandler)
	router.GET("/employees/", HomeHandler)
	router.POST("/employees/insert/", InsertHandler)
	router.POST("/employees/update/", UpdateHandler)
	router.GET("/employees/edit/:id", EditHandler)
	router.GET("/employees/delete/", DeleteHandler)
	router.NoRoute(NotFoundHandler)
	
	router.Run(":8000")
}

func HomeHandler(c *gin.Context){
	employees := domains.Employee{}.GetEmployees()

	c.HTML(http.StatusOK, "Home.html", gin.H{
		"employees" : employees,
	})
}

func EditHandler(c *gin.Context){
	employeeId := c.Param("id")
	employee := domains.Employee{}.GetEmployeeById(parseInt(employeeId))

	c.HTML(http.StatusOK, "EditEmployee.html", gin.H{
		"employee": employee,
	})
}

func InsertHandler(c *gin.Context){
	var employee domains.Employee

	employee.FirstName = c.PostForm("firstName")
	employee.LastName = c.PostForm("lastName")
	employee.Salary = parseInt(c.PostForm("salary")) * 100
	employee.Department = c.PostForm("department")

	domains.Employee.InsertEmployee(employee)
	
	c.Redirect(http.StatusFound, "/")
}

func UpdateHandler(c *gin.Context){
	var employee domains.Employee

	employee.Id = parseInt(c.PostForm("employeeId"))
	employee.Department = c.PostForm("department")
	employee.FirstName = c.PostForm("firstName")
	employee.LastName = c.PostForm("lastName")
	employee.Salary = parseInt(c.PostForm("salary")) * 100

	domains.Employee.UpdateEmployee(employee)

	c.Redirect(http.StatusFound, "/")
}

func DeleteHandler(c *gin.Context){
	var employeeId int
	employeeId = parseInt(c.Param("id"))

	employee := domains.Employee{}.GetEmployeeById(employeeId)
	domains.Employee.DeleteEmployee(employee)

	c.Redirect(http.StatusFound, "/")
}

func NotFoundHandler(c *gin.Context){
	c.HTML(http.StatusNotFound, "404.html", nil)
}

func parseInt(value string) (int){
	var convertedValue int

	fmt.Sscanf(value, "%d", &convertedValue)

	return convertedValue
}
