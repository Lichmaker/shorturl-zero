package apps

import "time"

type Apps struct {
	Id        int64     `gorm:"id"`
	AppId     string    `gorm:"app_id"`
	Name      string    `gorm:"name"`
	AppSecret string    `gorm:"app_secret"`
	CreatedAt time.Time `gorm:"created_at"`
}

// TableName 会将表名重写
func (Apps) TableName() string {
	return "apps"
}
