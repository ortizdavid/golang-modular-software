package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
	shared "github.com/ortizdavid/golang-modular-software/modules/shared/repositories"
)

type ProfessionalInfoRepository struct {
	db *database.Database
	*shared.BaseRepository[entities.ProfessionalInfo]
}

func NewProfessionalInfoRepository(db *database.Database) *ProfessionalInfoRepository {
	return &ProfessionalInfoRepository{
		db: db,
		BaseRepository: shared.NewBaseRepository[entities.ProfessionalInfo](db),
	}
}

func (repo *ProfessionalInfoRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.ProfessionalInfoData, error) {
	var professionalInfo entities.ProfessionalInfoData
	result := repo.db.WithContext(ctx).Table("employees.view_professional_info_data").Where("unique_id=?", uniqueId).First(&professionalInfo)
	return professionalInfo, result.Error
}

func (repo *ProfessionalInfoRepository) GetDataByEmployeeId(ctx context.Context, employeeId int64) (entities.ProfessionalInfoData, error) {
	var professionalInfo entities.ProfessionalInfoData
	result := repo.db.WithContext(ctx).Table("employees.view_professional_info_data").Where("employee_id=?", employeeId).First(&professionalInfo)
	return professionalInfo, result.Error
}

func (repo *ProfessionalInfoRepository) Exists(ctx context.Context, request entities.CreateProfessionalInfoRequest) (bool, error) {
	var professionalInfo entities.ProfessionalInfo
	result := repo.db.WithContext(ctx).
		Table("employees.professional_info").
		Where("employee_id=? AND department_id=? AND job_title_id=? AND employment_status_id=?", 
			request.EmployeeId,
			request.DepartmentId,
			request.JobTitleId,
			request.EmploymentStatusId).
		Find(&professionalInfo)
	return professionalInfo.EmployeeId != 0, result.Error
}
