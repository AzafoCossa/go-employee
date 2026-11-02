package domains

import (
	"database/sql"

	"azafocossa.com/employee/db"
)

type Employee struct{
	Id int
	FirstName string
	LastName string
	Salary int
	Department string
}

func (employee Employee) GetEmployees() []Employee{
	var database *sql.DB
	database = db.GetConnection()

	var error error

	var rows *sql.Rows

	rows, error = database.Query("SELECT * FROM Employee")

	if error != nil{
		panic(error.Error())
	}

	employee = Employee{}

	var employees []Employee
	employees = []Employee{}

	for rows.Next(){
		var id int
		var firstName string
		var lastName string
		var salary int
		var department string

		error = rows.Scan(&id, &firstName, &lastName, &salary, &department)

		if error != nil{
			panic(error.Error())
		}

		employee.Id = id
		employee.FirstName = firstName
		employee.LastName = lastName
		employee.Department = department
		employee.Salary = salary / 100

		employees = append(employees, employee)
	}

	defer database.Close()

	return employees
}

func (employee Employee)GetEmployeeById(employeeId int) Employee{
	var database *sql.DB
	database = db.GetConnection()

	var rows *sql.Rows
	rows, error := database.Query("SELECT * FROM Employee WHERE id = ?", employeeId)

	if(error != nil){
		panic(error.Error())
	}

	for rows.Next(){
		var id int
		var firstName string
		var lastName string
		var salary int
		var department string

		error = rows.Scan(&id, &firstName, &lastName, &salary, &department)

		if(error != nil){
			panic(error.Error())
		}

		employee.Department = department
		employee.FirstName = firstName
		employee.LastName = lastName
		employee.Salary = salary/100
		employee.Id = id
	}

	defer database.Close()

	return employee
}

func (employee Employee) InsertEmployee(){
	var database *sql.DB
	database = db.GetConnection()

	var error error

	var insert *sql.Stmt
	insert, error = database.Prepare("INSERT INTO Employee(firstName, lastName, salary, department) VALUES(?,?,?,?)")
	if error != nil{
		panic(error.Error())
	}

	insert.Exec(employee.FirstName, employee.LastName, employee.Salary, employee.Department)

	defer database.Close()
}

func (employee Employee) UpdateEmployee(){
	database := db.GetConnection()

	var update *sql.Stmt
	update, error := database.Prepare("UPDATE Employee SET firstName = ?, lastName = ?, salary = ?, department = ? WHERE id = ?")
	if error != nil{
		panic(error.Error())
	}

	update.Exec(employee.FirstName, employee.LastName, employee.Salary, employee.Department, employee.Id)

	defer database.Close()
}

func (employee Employee) DeleteEmployee(){
	var database *sql.DB
	database = db.GetConnection()
	
	var delete *sql.Stmt
	delete, error := database.Prepare("DELETE FROM Employee WHERE id = ?")

	if error != nil{
		panic(error.Error())
	}

	delete.Exec(employee.Id)
	defer database.Close()
}