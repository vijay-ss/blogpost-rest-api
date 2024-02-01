# blogpost-rest-api

A simple CRUD API for managing blog posts in a database.

## Tools
- Gin
- GORM
- ElephantSQL
- CompileDaemon
- Postman

Execute the following command for CompileDaemon to run:
- `CompileDaemon -command="./blogpostApi"`

## JWT Authentication

User is required to signup and login before making requests to blog post database. JWT with bcrypt password hashing is used.
