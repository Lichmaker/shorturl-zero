package shorts

import "gorm.io/gorm"

func GetByShort(db *gorm.DB, shortStr string) (Short, error) {
	model := Short{}
	result := db.Where(Short{
		Short: shortStr,
	}).First(&model)
	if result.Error == gorm.ErrRecordNotFound {
		return model, nil
	}
	return model, result.Error
}
