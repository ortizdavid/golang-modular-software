package services

import "github.com/ortizdavid/golang-modular-software/database"

type EmployeeService struct {
}

func NewEmployeeService(db *database.Database) *EmployeeService {
	return &EmployeeService{}
}