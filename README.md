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
- [x] Company
- [x] Reference
- [x] Employees
- [ ] Reports


## Architecture
![Modular Arquitecure](docs/_diagrams/Modular%20Architecture.png)


## Dependencies Between Modules
![Dependencies Between Modules](docs/_diagrams/Dependency%20Between%20Modules.png)

## Module Structure
![Structure](docs/_diagrams/Module%20Structure.png)


## Module Flow
![Flow](docs/_diagrams/Modules%20-%20Flow.png)


## Module Communication
![Comunication Between Modules](docs/_diagrams/Communication.png)


## Routing 
![Routing](docs/_diagrams/Routes.png)


## Database 
![Database](docs/_diagrams/Modular%20Database.png)

## Schema Objects 
![Schema Objects](docs/_diagrams/Schema-Objects.png)


## Module Creation
![Module Creation](docs/_diagrams/Module%20Creation.png)


## App Initialization
![App Initialization](docs/_diagrams/App%20Initalization.png)

## How to Run App
- Configure Database:
    - Open Postgres: ``psql -U postgres``
    - Create Database: ``CREATE DATABASE golang_modular_software;``
    - Add database to .env file -> var: DATABASE_MAIN_URL
- Run Application: ``go run main.go``
- User the url: http://localhost:4003
- Login with Default user:
    - Username: admin01
    - Password: 12345678