# Go Database Library

## Go Library Link

[Go Library DashLMS-Core-Database](https://pkg.go.dev/github.com/Dash-LMS/DashLMS-Core-Database/drivers/postgres)

## General Idea

This project provides a database operation library in Golang using the Echo framework. It supports multiple database drivers (MongoDB, PostgreSQL, and MySQL) with basic CRUD operations following CQRS principles.

## Technologies Used

- **Golang**: Core programming language.
- **Echo**: Web framework for routing and middleware.
- **GORM**: ORM for PostgreSQL and MySQL.
- **MongoDB Go Driver**: MongoDB driver for Go.
- **PowerShell**: Script for folder and file initialization.

## Project Structure

```sh
DashLMS-Core-Database/
│
├── interfaces/
│   └── database.go                         # Common database interface
│
├── drivers/
│   ├── mongo                               # MongoDB implementation
│   │   ├── mongo_command_driver.go     
│   │   └── mongo_query_driver.go      
│   ├── postgres                            # PostgreSQL implementation
│   │   ├── postgres_command_driver.go 
│   │   └── postgres_query_driver.go    
│   └── mysql                               # MySQL implementation
│       ├── mysql_command_driver.go 
│       └── mysql_query_driver.go        
│
├── factory/
│   └── database_factory.go                 # Database driver selection logic
│
├── utils/
│   └── query_validator.go                  # Query validation and security checks
│
├── tests/                                  # Separate test folder
│   ├── mongo/
│   │   ├── mongo_command_driver_test.go
│   │   └── mongo_query_driver_test.go
│   ├── postgres/
│   │   ├── postgres_command_driver_test.go
│   │   └── postgres_query_driver_test.go
│   ├── mysql/
│   │   ├── mysql_command_driver_test.go
│   │   └── mysql_query_driver_test.go
│   ├── factory/
│   │   └── database_factory_test.go
│   └── utils/
│       └── query_validator_test.go
│
├── go.mod                                  # Go module configuration
├── README.md                               # Library documentation
├── LICENSE.md                              # Library license
└── .gitignore                              # .gitignore
```

## Content Table

- **Interfaces**: Common interface for database operations.
- **Drivers**: Database-specific implementations.
- **Factory**: Logic for driver selection based on the configuration.
- **Utils**: Utility functions for query validation.
- **Initialization**: Scripts for file structure creation.

## Initialization Steps

1. Open PowerShell in the project directory.
2. Run the initialization script:

   ```powershell
   .\init.ps1
   ```

3. Initialize the Go module:

   ```sh
   go mod tidy
   ```

4. Import the necessary drivers and use the library as needed.

## Usage

```go
package main

import (
  "fmt"
  "dashlms.com/dashlms-core-database/factory"
)

func main() {
  db, err := factory.NewDatabase("mongo")
  if err != nil {
    panic(err)
  }
  err = db.Connect("mongodb://localhost:27017", "mydb")
  if err != nil {
    panic(err)
  }
  defer db.Close()

  fmt.Println("Database connected successfully")
}
```

## Commit Rules

The version increments are based on commit messages

- #major → Increments major version (e.g., v1.0.0 → v2.0.0)
- #minor → Increments minor version (e.g., v1.0.0 → v1.1.0)
- Default → Increments patch version (e.g., v1.0.0 → v1.0.1)
