package services

import (
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/repositories"
)

type ProjectService struct {
	repository *repositories.ProjectRepository
}

func NewProjectService(db *database.Database) *ProjectService {
	return &ProjectService{
		repository: repositories.NewProjectRepository(db),
	}
}