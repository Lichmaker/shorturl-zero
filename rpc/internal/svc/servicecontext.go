package svc

import (
	"time"

	"github.com/lichmaker/short-url-micro/rpc/internal/config"
	mysqldriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	GormDB *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysqldriver.New(mysqldriver.Config{
		DSN: c.Mysql.DataSource,
	}))
	if err != nil {
		panic(err)
	}
	dbCore, err := db.DB()
	if err != nil {
		panic(err)
	}
	dbCore.SetConnMaxLifetime(time.Duration(c.Mysql.MaxLifeTime))
	dbCore.SetMaxIdleConns(int(c.Mysql.MaxIdleConn))
	dbCore.SetMaxOpenConns(int(c.Mysql.MaxOpenConn))

	return &ServiceContext{
		Config: c,
		GormDB: db,
	}
}
