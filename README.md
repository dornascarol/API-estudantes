# Students API 👩‍🎓 👨‍🎓
<br>

_Project I learned in the Golang do Zero course_

<h2> Introduction </h2>
The objective of the application is to create a system to control registered and active students on the course platform. Based on GET, POST, PUT and DELETE requests to register, list, update and delete profiles. Using the SQLite database.

## Technologies used
* VS Code
* Golang (Go)
* Insomnia
* Echo v4
* SQLite
* GORM
* Zerolog
* Swagger

## Tools
Go extension installed in VS Code with version 0.41.2

SQLite extension installed in VS Code with version 0.14.1

SQLite Viewer extension installed in VS Code with version 0.5.8

Installed C/C++ extension in VS Code with version 1.20.5

Makefile Tools extension installed in VS Code with version 0.9.10

The program <a href="https://insomnia.rest/download" target="_blank" > Insomnia</a> was used to test route requests by simulating the Front-end.

The web framework <a href="https://github.com/labstack/echo" target="_blank"> ECHO</a> version 4 was used.

The object-relational mapping (ORM) tool <a href="https://gorm.io/docs/connecting_to_the_database.html" target="_blank"> GORM</a> was used to connect to the SQLite database.

The <a href="https://github.com/rs/zerolog" target="_blank"> Zerolog</a> package was used.

<a href="https://github.com/swaggo/echo-swagger" target="_blank"> Swagger</a> was used with Echo.

## Running the project
- Command to run the server:
```
go run main.go
```

- Stop running the server: in the terminal, click the "Ctrl" and "C" keys.

- Command to initialize the package manager (_go.mod_):
```
go mod init
```

- Command to check, update and download modules in the package manager (_go.mod_):
```
go mod tidy
```

- After the Makefile is configured, the new command to run the server:
```
make run
```

## Endpoints

| Method | URL             | Description                                 |
| ------ | --------------  | --------------------------------------------|
| GET    | /estudantes     | List all students                           |
| POST   | /estudantes     | Register a student                          |
| GET    | /estudantes/:id | Get information about a specific student    |
| PUT    | /estudantes/:id | Update information for a specific student   |
| DELETE | /estudantes/:id | Delete a specific student                   |


## Student structure
- Nome
- CPF
- Email
- Idade
- Ativo

## Swagger
From the <a href="https://github.com/swaggo/echo-swagger" target="_blank"> documentation </a>, follow these steps:

- Commands to download the package:
```
go get -d github.com/swaggo/swag/cmd/swag
```
```
go install github.com/swaggo/swag/cmd/swag@latest
```

- Command to initialize:
```
swag init
```

- Command to download the echo-swagger lib:
```
go get -u github.com/swaggo/echo-swagger
```

- Import into the _api.go_ file following my user in this repository:
```
echoSwagger "github.com/swaggo/echo-swagger"

_ "github.com/dornascarol/API-estudantes/docs"
```

- Command to update in the package manager (_go.mod_):
```
go mod tidy
```

- Command to run the server:
```            
go run main.go
```

- To test Swagger, type the following URL into your browser:
```
http://localhost:8080/swagger/index.html
```

## Project status
:heavy_check_mark: Application completed.
