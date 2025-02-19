package mysql

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bhtoan2204/user/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetPool() {
	db, err := global.MDB.DB()
	if err != nil {
		// global.Logger.Error("Failed to set pool", zap.Error(err))
		panic(err)
	}
	db.SetMaxOpenConns(global.Config.MySQLConfig.MaxOpenConns)
	db.SetMaxIdleConns(global.Config.MySQLConfig.MaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(global.Config.MySQLConfig.MaxLifetime) * time.Second)
}

func NewDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		global.Config.MySQLConfig.User,
		global.Config.MySQLConfig.Pass,
		global.Config.MySQLConfig.Host,
		strconv.Itoa(global.Config.MySQLConfig.Port),
		global.Config.MySQLConfig.Name,
		global.Config.MySQLConfig.Charset,
		global.Config.MySQLConfig.ParseTime,
		global.Config.MySQLConfig.Loc,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	global.MDB = db
	SetPool()
	return nil
}
