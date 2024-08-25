package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
)

type FeatureRepository struct {
	db *database.Database
}

func NewFeatureRepository(db *database.Database) *FeatureRepository {
	return &FeatureRepository{
		db: db,
	}
}

func (repo *FeatureRepository) Create(ctx context.Context, company entities.Feature) error {
	result := repo.db.WithContext(ctx).Create(&company)
	return result.Error
}

func (repo *FeatureRepository) Update(ctx context.Context, company entities.Feature) error {
	result := repo.db.WithContext(ctx).Save(&company)
	return result.Error
}

func (repo *FeatureRepository) Delete(ctx context.Context, company entities.Feature) error {
	result := repo.db.WithContext(ctx).Delete(&company)
	return result.Error
}

func (repo *FeatureRepository) FindAll(ctx context.Context) ([]entities.Feature, error) {
	var features []entities.Feature
	result := repo.db.WithContext(ctx).Find(&features)
	return features, result.Error
}

func (repo *FeatureRepository) FindAllLimit(ctx context.Context, limit int, offset int) ([]entities.FeatureData, error) {
	var features []entities.FeatureData
	result := repo.db.WithContext(ctx).Table("configurations.view_feature_data").Limit(limit).Offset(offset).Find(&features)
	return features, result.Error
}

func (repo *FeatureRepository) FindById(ctx context.Context, id int) (entities.Feature, error) {
	var company entities.Feature
	result := repo.db.WithContext(ctx).First(&company, id)
	return company, result.Error
}

func (repo *FeatureRepository) FindByUniqueId(ctx context.Context, uniqueId string) (entities.Feature, error) {
	var company entities.Feature
	result := repo.db.WithContext(ctx).Where("unique_id=?", uniqueId).First(&company)
	return company, result.Error
}

func (repo *FeatureRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.FeatureData, error) {
	var company entities.FeatureData
	result := repo.db.WithContext(ctx).Table("configurations.view_feature_data").Where("unique_id=?", uniqueId).First(&company)
	return company, result.Error
}

func (repo *FeatureRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("configurations.features").Count(&count)
	return count, result.Error
}

func (repo *FeatureRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.FeatureData, error) {
	var features []entities.FeatureData
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM configurations.view_feature_data WHERE feature_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&features)
	return features, result.Error
}

func (repo *FeatureRepository) CountByParam(ctx context.Context, param string) (int64, error) {
	var count int64
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM configurations.view_feature_data WHERE feature_name LIKE ? OR code LIKE ?", likeParam, likeParam).
		Scan(&count)
	return count, result.Error
}

func (repo *FeatureRepository) ExistsByName(ctx context.Context, featureName string) (bool, error) {
	var feature entities.Feature
	result := repo.db.WithContext(ctx).Where("feature_name=?", featureName).Find(&feature)
	return feature.FeatureId != 0, result.Error
}

func (repo *FeatureRepository) ExistsByModule(ctx context.Context, moduleId int, featureName string) (bool, error) {
	var feature entities.Feature
	result := repo.db.WithContext(ctx).Where("module_id=? AND feature_name=?", moduleId, featureName).Find(&feature)
	return feature.FeatureId != 0, result.Error
}
