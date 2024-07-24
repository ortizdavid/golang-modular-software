package services

import (
	"context"
	"fmt"

	"github.com/ortizdavid/golang-modular-software/modules/authentication/repositories"
	"gorm.io/gorm"
)

type UserService struct {
	repository *repositories.UserRepository
	roleRepository *repositories.RoleRepository
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		repository: repositories.NewUserRepository(db),
		roleRepository: repositories.NewRoleRepository(db),
	}
}

func (s *UserService) UserHasRoles(ctx context.Context, userId int64, comparedRoles ...string) (bool, error) {
	for _, role := range comparedRoles {
		exists, err := s.roleRepository.ExistsByCode(ctx, role)
		if err != nil {
			return false, fmt.Errorf("error validating role '%s': %w", role, err)
		}
		if !exists {
			return false, fmt.Errorf("role '%s' does not exist", role)
		}
	}
	hasRole, err := s.repository.HasRoles(ctx, userId, comparedRoles)
	if err != nil {
		return false, fmt.Errorf("error checking roles for user ID %d: %w", userId, err)
	}
	return hasRole, nil
}