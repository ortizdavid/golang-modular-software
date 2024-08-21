package services

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/company/entities"
	"github.com/ortizdavid/golang-modular-software/modules/company/repositories"
)

type StatisticsService struct {
	branchRepository     *repositories.BranchRepository
	officeRepository     *repositories.OfficeRepository
	departmentRepository *repositories.DepartmentRepository
	roomRepository       *repositories.RoomRepository
	projectRepository    *repositories.ProjectRepository
	policyRepository     *repositories.PolicyRepository
}

func NewStatisticsService(db *database.Database) *StatisticsService {
	return &StatisticsService{
		branchRepository:    repositories.NewBranchRepository(db),
		officeRepository:    repositories.NewOfficeRepository(db),
		departmentRepository: repositories.NewDepartmentRepository(db),
		roomRepository:      repositories.NewRoomRepository(db),
		projectRepository:   repositories.NewProjectRepository(db),
		policyRepository:    repositories.NewPolicyRepository(db),
	}
}

func (s *StatisticsService) GetStatistics(ctx context.Context) (entities.Statistics, error) {
	branches, err := s.branches(ctx)
	if err != nil {
		return entities.Statistics{}, err
	}
	offices, err := s.offices(ctx)
	if err != nil {
		return entities.Statistics{}, err
	}
	departments, err := s.departments(ctx)
	if err != nil {
		return entities.Statistics{}, err
	}
	rooms, err := s.rooms(ctx)
	if err != nil {
		return entities.Statistics{}, err
	}
	projects, err := s.projects(ctx)
	if err != nil {
		return entities.Statistics{}, err
	}
	policies, err := s.policies(ctx)
	if err != nil {
		return entities.Statistics{}, err
	}
	return entities.Statistics{
		Branches:     branches,
		Offices:      offices,
		Departments:  departments,
		Rooms:        rooms,
		Projects:     projects,
		Policies:     policies,
	}, nil
}

func (s *StatisticsService) branches(ctx context.Context) (int64, error) {
	return s.branchRepository.Count(ctx)
}

func (s *StatisticsService) offices(ctx context.Context) (int64, error) {
	return s.officeRepository.Count(ctx)
}

func (s *StatisticsService) departments(ctx context.Context) (int64, error) {
	return s.departmentRepository.Count(ctx)
}

func (s *StatisticsService) rooms(ctx context.Context) (int64, error) {
	return s.roomRepository.Count(ctx)
}

func (s *StatisticsService) projects(ctx context.Context) (int64, error) {
	return s.projectRepository.Count(ctx)
}

func (s *StatisticsService) policies(ctx context.Context) (int64, error) {
	return s.policyRepository.Count(ctx)
}
