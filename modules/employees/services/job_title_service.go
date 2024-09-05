package services

import "github.com/ortizdavid/golang-modular-software/database"

type JobTitleService struct {
}

func NewJobTitleService(db *database.Database) *JobTitleService {
	return &JobTitleService{}
}