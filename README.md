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
- [ ] Company
- [ ] Employees
- [ ] Reference


## Architecture
![Modular Arquitecure](docs/Modular-Architecture.png)


## Module Structure
![Structure](docs/Module%20Structure.png)


## Module Comunication
![Comunication Between Modules](docs/Comunication.jpg)


## Routing 
![Routing](docs/Routes.jpg)


## Database 
![Database Schemas](docs/Modular%20Database.jpg)
![Schema Objects](docs/Schema-Objects.png)

## Module Creation
![Module Creation](docs/Module%20Creation.png)

## How to Run App
- Configure Database:
    - Open Postgres: ``psql -U postgres``
    - Create Database: ``CREATE DATABASE golang_modular_software;``
- Run Application: ``go run main.go``
- User the url: http://localhost:4003
- Login with Default user:
    - Username: admin01
    - Password: 12345678