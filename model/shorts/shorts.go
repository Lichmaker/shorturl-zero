package shorts

import "time"

type Short struct {
	Id        int64     `gorm:"id" json:"id"`
	Long      string    `gorm:"long" json:"long"`
	Short     string    `gorm:"short" json:"short"`
	ExpiredAt int64     `gorm:"expired_at" json:"expired_at"`
	CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"`
}

// type Tabler interface {
// 	TableName() string
// }

// TableName 会将表名重写
func (Short) TableName() string {
	return "shorts"
}
