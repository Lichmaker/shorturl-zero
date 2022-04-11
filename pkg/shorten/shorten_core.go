package shorten

import (
	"github.com/lichmaker/short-url-micro/model/shorts"
	"github.com/lichmaker/short-url-micro/pkg/hash"
	"github.com/lichmaker/short-url-micro/pkg/helpers"
	"gorm.io/gorm"
)

func Make(db *gorm.DB, long string) (shorts.Short, error) {
	// 补充协议头
	long = helpers.FillHttpScheme(long)

	// 哈希
	short := hash.Make(long)

	// 查库
	shortModel, err := shorts.GetByShort(db, short)
	if err != nil {
		return shortModel, err
	}
	if shortModel.Id > 0 {
		return shortModel, nil
	}

	// 写库
	shortModel = shorts.Short{
		Long:      long,
		Short:     short,
		ExpiredAt: 0,
	}
	res := db.Create(&shortModel)
	if res.Error != nil {
		return shortModel, res.Error
	}
	return shortModel, nil
}
