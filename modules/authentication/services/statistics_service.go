package services

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	"github.com/ortizdavid/golang-modular-software/modules/authentication/repositories"
)

type StatisticsService struct {
	userRepository        *repositories.UserRepository
	roleRepository        *repositories.RoleRepository
	permissionRepository  *repositories.PermissionRepository
	loginRepository       *repositories.LoginActivityRepository
}

func NewStatisticsService(db *database.Database) *StatisticsService {
	return &StatisticsService{
		userRepository:        repositories.NewUserRepository(db),
		roleRepository:        repositories.NewRoleRepository(db),
		permissionRepository:  repositories.NewPermissionRepository(db),
		loginRepository:       repositories.NewLoginActivityRepository(db),
	}
}

func (s *StatisticsService) GetStatistics(ctx context.Context)  (entities.Statistics, error) {
	users, err := s.users(ctx)
	if err != nil {
		return entities.Statistics{}, err
	}
	activeUsers, err := s.activeUsers(ctx)
	if err != nil {
		return entities.Statistics{}, err
	}
	inactiveUsers, err := s.inactiveUsers(ctx)
	if err != nil {
		return entities.Statistics{}, err
	}
	onlineUsers, err := s.onlineUsers(ctx)
	if err != nil {
		return entities.Statistics{}, err
	}
	offlineUsers, err := s.offlineUsers(ctx)
	if err != nil {
		return entities.Statistics{}, err
	}
	roles, err := s.roles(ctx)
	if err != nil {
		return entities.Statistics{}, err
	}
	permissions, err := s.permissions(ctx)
	if err != nil {
		return entities.Statistics{}, err
	}
	loginActivity, err := s.loginActivity(ctx)
	if err != nil {
		return entities.Statistics{}, err
	}
	return entities.Statistics{
		Users:         users,
		ActiveUsers:   activeUsers,
		InactiveUsers: inactiveUsers,
		OnlineUsers:   onlineUsers,
		OfflineUsers:  offlineUsers,
		Roles:         roles,
		Permissions:   permissions,
		LoginActivity: loginActivity,
	}, nil
}

func (s *StatisticsService) users(ctx context.Context) (int64, error) {
	return s.userRepository.Count(ctx)
}

func (s *StatisticsService) activeUsers(ctx context.Context) (int64, error) {
	return s.userRepository.CountByStatus(ctx, true)
}

func (s *StatisticsService) inactiveUsers(ctx context.Context) (int64, error) {
	return s.userRepository.CountByStatus(ctx, false)
}

func (s *StatisticsService) roles(ctx context.Context) (int64, error) {
	return s.roleRepository.Count(ctx)
}

func (s *StatisticsService) permissions(ctx context.Context) (int64, error) {
	return s.permissionRepository.Count(ctx)
}

func (s *StatisticsService) loginActivity(ctx context.Context) (int64, error) {
	sumLogin, sumLogout, err := s.loginRepository.SumLoginAndLogout(ctx)
	return (sumLogin+sumLogout), err
}

func (s *StatisticsService) onlineUsers(ctx context.Context) (int64, error) {
	return s.loginRepository.CountByStatus(ctx, string(entities.ActivityStatusOnline))
}

func (s *StatisticsService) offlineUsers(ctx context.Context) (int64, error) {
	return s.loginRepository.CountByStatus(ctx, string(entities.ActivityStatusOffline))
}
