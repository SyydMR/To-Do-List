# To-Do List API

A RESTful API for managing to-do tasks built using Golang. This project uses JWT for authentication, GORM as an ORM library, and MySQL as the database. 

## Features

- **User Authentication**: Secure user authentication using JSON Web Tokens (JWT).
- **CRUD Operations**: Create, read, update, and delete to-do tasks.
- **Data Persistence**: Store user and task information in a MySQL database.
- **ORM**: Easy data handling and migration through GORM.

## Tech Stack

- **Golang**: Core programming language for the project.
- **JWT**: JSON Web Token for secure user authentication.
- **GORM**: Golang ORM library for handling MySQL operations.
- **MySQL**: Database for storing user and task data.

### Authentication

- **POST** `/register`: Register a new user
- **POST** `/login`: Authenticate and retrieve a JWT

### To-Do Tasks

- **GET** `/items`: Retrieve all tasks for the authenticated user
- **POST** `/items`: Create a new task
- **GET** `/items/{id}`: Retrieve a specific task by ID
- **PUT** `/items/{id}`: Update a task by ID
- **DELETE** `/items/{id}`: Delete a task by ID
- **POST** `/items/{id}`: Mark a specific task as complete by its ID

## Example Usage

Here’s how to use the API with `curl`:

### Register a User

```bash
curl -X POST http://localhost:8090/register -d '{"username": "yourUsername", "password": "yourPassword"}'
```

### Login

```bash
curl -X POST http://localhost:8090/login -d '{"username": "yourUsername", "password": "yourPassword"}'
```

This will return a JWT token, which is needed for all subsequent requests.

### Get All Tasks

```bash
curl -X GET http://localhost:8090/items -H "Authorization: Bearer YOUR_JWT_TOKEN"
```
