# Golang Modular Software Example
Example of a complete modular sofware, written in Golang 


## Status
- In progress


## Tools
- Golang
- Fiber Framework
- PostgreSQL


## Modules
- [x] Configurations 
- [x] Authentication
- [ ] Human Resources
- [ ] Customers

## Module Comunication
![Comunication Between Modules](docs/Comunication.jpg)

## Routing 
![Routing](docs/Routes.jpg)

## Database 
![Database Schemas](docs/Database.jpg)
![Schema Objects](docs/Schema-Objects.png)


## How to Run App
- Configure Database:
    - Copy all database scripts placed in [datatabase dir](database)
- Run Application: ``go run main.go``
- User the url: http://localhost:9000
- Login with Default user:
    - Username: admin@user.com
    - Password: 12345678