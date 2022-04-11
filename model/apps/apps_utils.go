package apps

import (
	"gorm.io/gorm"
)

func FindByAppId(db *gorm.DB, appId string) (Apps, error) {
	model := Apps{}
	result := db.Where(Apps{
		AppId: appId,
	}).First(&model)
	if result.Error == gorm.ErrRecordNotFound {
		return model, nil
	}
	return model, result.Error
}
