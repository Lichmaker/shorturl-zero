package statistics

import (
	"time"
)

type Statistics struct {
	Id        int64     `gorm:"id" json:"id"`
	Host      string    `gorm:"host" json:"host"`
	Short     string    `gorm:"short" json:"short"`
	Long      string    `gorm:"long" json:"long"`
	Counter   int64     `gorm:"counter" json:"counter"`
	CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"`
}

// TableName 会将表名重写
func (Statistics) TableName() string {
	return "statistics"
}
