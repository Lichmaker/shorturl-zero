package statistics

import (
	"net/url"
	"time"

	"github.com/lichmaker/short-url-micro/model/shorts"
	"github.com/lichmaker/short-url-micro/model/statistics"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

func Do(db *gorm.DB, m shorts.Short) {
	var stat statistics.Statistics
	r := db.Where(statistics.Statistics{
		Short: m.Short,
	}).First(&stat)
	if r.Error == gorm.ErrRecordNotFound {
		parseUrl, err := url.Parse(m.Long)
		if err != nil {
			logx.Infof("url解析失败，不进行统计 : %s", m.Long)
			return
		}

		stat = statistics.Statistics{
			Host:    parseUrl.Host,
			Short:   m.Short,
			Long:    m.Long,
			Counter: 1,
		}
		result := db.Create(&stat)
		if result.Error != nil {
			logx.Error("失败：" + result.Error.Error())
		}
	} else {
		db.Exec("UPDATE `statistics` SET `counter` = `counter` + 1, `updated_at` = ? WHERE `short` = ?;", time.Now(), m.Short)
	}
}
