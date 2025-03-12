# ToDo API

## ¿What is ToDo API?
ToDo API is a simple RESTful application built with Go that allows you to manage a list of tasks. It supports creating tasks, retrieving tasks, updating task details, and deleting tasks (commonly known as CRUD operations).

The application uses an in-memory data store to manage tasks and follows a clean architecture with separate layers for handlers, services, and repositories. It also includes unit tests to ensure reliability.

---
## Requirements

- Go 1.20 or higher
- Chi router (github.com/go-chi/chi/v5)
- Testing libraries (github.com/stretchr/testify and github.com/golang/mock/gomock for mocking)

## Installation

#### 1. Clone project: 

```bash
https://github.com/castellarmartinez/todo-list-go.git
```

#### 2. Install dependencies

Run the following command to download and install the required dependencies:

```
go mod download
```

#### 3. Build the application

Build the application using the following command:

```
go build -o todo-api
```

#### 4. How to run the app

Run the application using the following command:

```
go run main.go
```

## How to run the tests

To run all tests in the project, use the following command:

```
go test -v ./...
```

To check test coverage, generate a coverage profile and display it:

```
go test --coverprofile=coverage.out ./...
```

```
go tool cover --func=coverage.out
```

## How to use the API

The server is running on: [http://localhost:8080](http://localhost:8080) (port 8080 is hardcoded).Below are the available endpoints:

#### 1. Create new task:

To create a new tasks a POST http request can be sent to: [http://localhost:8080/tasks](http://localhost:8080/tasks):

```
curl -i -X POST http://localhost:8080/tasks -H "Content-Type: application/json" -d "{\"title\": \"Peluquearme\", \"description\": \"Necesito peluquearme\"}"
```

The description is not mandatory:
```
curl -i -X POST http://localhost:8080/tasks -H "Content-Type: application/json" -d "{\"title\": \"Peluquearme\"}"
```

##### Example response
```
{
    "id": 3,
    "title": "New Task",
    "description": "This is a new task",
    "completed": false
}
```

##### Example error response
```
{
    "error": "Title is required"
}
```

#### 2. Get all tasks

To get all tasks a GET http request can be sent to: [http://localhost:8080/tasks](http://localhost:8080/tasks):

```
curl -i -X GET http://localhost:8080/tasks
```

##### Example response
```

[
  {
    "id": 3,
    "title": "New Task",
    "description": "This is a new task",
    "completed": false
  },
  {
    "id": 4,
    "title": "Old Task",
    "description": "This is an old task",
    "completed": true
  }
]
```


#### 3. Get task by id

To get a task by its id an GET http request can be sent to: [http://localhost:8080/tasks/:id](http://localhost:8080/tasks/:id):

```
curl -i -X GET http://localhost:8080/tasks/1
```

##### Example response
```
{
    "id": 3,
    "title": "New Task",
    "description": "This is a new task",
    "completed": false
}
```

##### Example error response
```
{
    "error": "The task with id 999 does not exist"
}
```

#### 4. Update task by id

To update a task by its id a PUT http request can be sent to: [http://localhost:8080/tasks/:id](http://localhost:8080/tasks/:id):

```
curl -i -X PUT http://localhost:8080/tasks/1 -H "Content-Type: application/json" -d "{\"completed\": true}"
```

NOTE: The title, description and completed values can be updated, sending a body with only one field or with multiple fields

##### Example response
```
{
    "id": 1,
    "title": "New Task",
    "description": "This is a new task",
    "completed": true
}
```

##### Example error response
```
{
    "error": "The task with id 999 does not exist"
}
```

#### 5. Delete id

To delete a task by its id a DELETE http request can be sent to: [http://localhost:8080/tasks/:id](http://localhost:8080/tasks/:id):

```
curl -i -X DELETE http://localhost:8080/tasks/1
```

##### Example error response
```
{
    "error": "The task with id 999 does not exist"
}
```

## Construido con: 

- [Go](https://go.dev/) - A statically typed, compiled programming language.
- [Chi](https://github.com/go-chi/chi) - A lightweight, idiomatic, and composable router for building Go HTTP services.
- [Testify](https://github.com/stretchr/testify) - A toolkit for writing unit tests in Go.

---
## Autor 
**David Castellar Martínez** [[GitHub](https://github.com/castellarmartinez/)]
[[LinkedIn](https://www.linkedin.com/in/castellarmartinez/)]

