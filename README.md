# Simple Employee Web System

A small employee management web application built with Go, HTML templates,
Bootstrap, and MySQL.

## Requirements

- Go 1.27.3 or a compatible newer Go version
- MySQL Server
- A terminal opened in the project root

## Install dependencies

Go dependencies are declared in `go.mod`. Download them with:

```bash
go mod download
```

You can also verify and clean the dependency files with:

```bash
go mod tidy
```

## Configure MySQL

The application connects to MySQL using the settings in
[`db/DBConnection.go`](db/DBConnection.go). Before starting the application,
make sure MySQL is running and create the database and table:

```sql
CREATE DATABASE go_employee;

USE go_employee;

CREATE TABLE Employee (
    id INT AUTO_INCREMENT PRIMARY KEY,
    firstName VARCHAR(50) NOT NULL,
    lastName VARCHAR(50) NOT NULL,
    salary INT NOT NULL,
    department VARCHAR(75) NOT NULL
);
```

Update the MySQL username, password, and database name in
[`db/DBConnection.go`](db/DBConnection.go) to match your local setup. Do not
commit real credentials to source control.

## Run the application

From the project root, run:

```bash
go run .
```

The server listens on port `8000`. Open
[http://localhost:8000](http://localhost:8000) in a browser.

Because the application loads templates and static assets using relative paths,
run the command from the project root.

## Build and run

To create an executable:

```bash
go build -o employee
./employee
```

Stop the server with `Ctrl+C`.
