package repositories

import (
	"context"

	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/configurations/entities"
)

type FeatureFlagRepository struct {
	db *database.Database
}

func NewFeatureFlagRepository(db *database.Database) *FeatureFlagRepository {
	return &FeatureFlagRepository{
		db: db,
	}
}

func (repo *FeatureFlagRepository) Create(ctx context.Context, featureFlag entities.FeatureFlag) error {
	result := repo.db.WithContext(ctx).Create(&featureFlag)
	return result.Error
}

func (repo *FeatureFlagRepository) Update(ctx context.Context, featureFlag entities.FeatureFlag) error {
	result := repo.db.WithContext(ctx).Save(&featureFlag)
	return result.Error
}

func (repo *FeatureFlagRepository) Delete(ctx context.Context, featureFlag entities.FeatureFlag) error {
	result := repo.db.WithContext(ctx).Delete(&featureFlag)
	return result.Error
}

func (repo *FeatureFlagRepository) FindAll(ctx context.Context) ([]entities.FeatureFlag, error) {
	var features []entities.FeatureFlag
	result := repo.db.WithContext(ctx).Find(&features)
	return features, result.Error
}

func (repo *FeatureFlagRepository) FindAllLimit(ctx context.Context, limit int, offset int) ([]entities.FeatureFlagData, error) {
	var features []entities.FeatureFlagData
	result := repo.db.WithContext(ctx).Table("configurations.view_feature_data").Limit(limit).Offset(offset).Find(&features)
	return features, result.Error
}

func (repo *FeatureFlagRepository) FindById(ctx context.Context, id int) (entities.FeatureFlag, error) {
	var featureFlag entities.FeatureFlag
	result := repo.db.WithContext(ctx).First(&featureFlag, id)
	return featureFlag, result.Error
}

func (repo *FeatureFlagRepository) FindByUniqueId(ctx context.Context, uniqueId string) (entities.FeatureFlag, error) {
	var featureFlag entities.FeatureFlag
	result := repo.db.WithContext(ctx).Where("unique_id=?", uniqueId).First(&featureFlag)
	return featureFlag, result.Error
}

func (repo *FeatureFlagRepository) GetDataByUniqueId(ctx context.Context, uniqueId string) (entities.FeatureFlagData, error) {
	var featureFlag entities.FeatureFlagData
	result := repo.db.WithContext(ctx).Table("configurations.view_feature_data").Where("unique_id=?", uniqueId).First(&featureFlag)
	return featureFlag, result.Error
}

func (repo *FeatureFlagRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := repo.db.WithContext(ctx).Table("configurations.features").Count(&count)
	return count, result.Error
}

func (repo *FeatureFlagRepository) Search(ctx context.Context, param string, limit int, offset int) ([]entities.FeatureFlagData, error) {
	var features []entities.FeatureFlagData
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT * FROM configurations.view_feature_data WHERE feature_name LIKE ? OR email LIKE ?", likeParam, likeParam).
		Limit(limit).
		Offset(offset).
		Scan(&features)
	return features, result.Error
}

func (repo *FeatureFlagRepository) CountByParam(ctx context.Context, param string) (int64, error) {
	var count int64
	likeParam := "%" + param + "%"
	result := repo.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM configurations.view_feature_data WHERE feature_name LIKE ? OR email LIKE ?", likeParam, likeParam).
		Scan(&count)
	return count, result.Error
}

func (repo *FeatureFlagRepository) ExistsByFeatureId(ctx context.Context, featureId int) (bool, error) {
	var feature entities.FeatureFlag
	result := repo.db.WithContext(ctx).Where("feature_id=?", featureId).Find(&feature)
	return feature.FlagId != 0, result.Error
}
