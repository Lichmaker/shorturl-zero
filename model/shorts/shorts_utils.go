package shorts

import (
	"context"

	"gorm.io/gorm"
)

func GetByShort(ctx context.Context, db *gorm.DB, shortStr string) (Short, error) {
	model := Short{}
	result := db.WithContext(ctx).Where(Short{
		Short: shortStr,
	}).First(&model)
	if result.Error == gorm.ErrRecordNotFound {
		return model, nil
	}
	return model, result.Error
}
