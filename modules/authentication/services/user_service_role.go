package services

import (
	"context"
	"fmt"
	"time"

	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/golang-modular-software/common/apperrors"
	entities "github.com/ortizdavid/golang-modular-software/modules/authentication/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/entities"
)

func (s *UserService) AssignUserRole(ctx context.Context, userId int64, request entities.AssignUserRoleRequest) error {
	if err := request.Validate(); err != nil {
		return apperrors.NewBadRequestError(err.Error())
	}
	user, err := s.repository.FindById(ctx, userId)
	if err != nil {
		return apperrors.NewNotFoundError("user not found. invalid user id")
	}
	role, err := s.roleRepository.FindById(ctx, request.RoleId)
	if err != nil {
		return apperrors.NewNotFoundError("role not found. invalid role id")
	}
	exists, err := s.userRoleRepository.ExistsByUserAndRole(ctx, userId, request.RoleId)
	if err != nil {
		return err
	}
	if exists {
		return apperrors.NewConflictError(fmt.Sprintf("role '%s' already assigned to user '%s'", role.RoleName, user.UserName))
	}
	userRole := entities.UserRole{
		UserId: userId,
		RoleId: request.RoleId,
		BaseEntity: shared.BaseEntity{
			UniqueId:  encryption.GenerateUUID(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	err = s.userRoleRepository.Create(ctx, userRole)
	if err != nil {
		return apperrors.NewInternalServerError("error while assigning role: " + err.Error())
	}
	return nil
}

func (s *UserService) AssociateUserToRole(ctx context.Context, request entities.AssociateUserRequest) error {
	user, err := s.repository.FindById(ctx, request.UserId)
	if err != nil {
		return apperrors.NewNotFoundError("user not found. invalid user id")
	}
	exists, err := s.userAssociationRepository.Exists(ctx, request.EntityId)
	if err != nil {
		return err
	}
	if exists {
		return apperrors.NewConflictError("entity already associated to an user.")
	}
	userAssociation := entities.UserAssociation{
		UserId:     user.UserId,
		EntityId:   request.EntityId,
		EntityName: request.EntityName,
		BaseEntity: shared.BaseEntity{
			UniqueId:  encryption.GenerateUUID(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	err = s.userAssociationRepository.Create(ctx, userAssociation)
	if err != nil {
		return apperrors.NewInternalServerError("error while associating user to a role: " + err.Error())
	}
	return nil
}

func (s *UserService) RemoveUserRole(ctx context.Context, uniqueId string) error {
	userRole, err := s.userRoleRepository.FindByUniqueId(ctx, uniqueId)
	if err != nil {
		return apperrors.NewNotFoundError("user role not found. invalid unique id")
	}
	err = s.userRoleRepository.Delete(ctx, userRole)
	if err != nil {
		return apperrors.NewInternalServerError("error while removing role: " + err.Error())
	}
	return nil
}
