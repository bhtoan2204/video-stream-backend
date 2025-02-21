package initialize

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bhtoan2204/user/global"
	"github.com/bhtoan2204/user/internal/infrastructure/db/mysql/persistent_object"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func migrateTables() {
	models := []interface{}{
		&persistent_object.User{},
		&persistent_object.Permission{},
		&persistent_object.Role{},
		&persistent_object.UserSettings{},
	}
	err := global.MDB.AutoMigrate(models...)
	if err != nil {
		global.Logger.Error("Failed to migrate tables", zap.Error(err))
		panic(err)
	}
}

func setPool() {
	db, err := global.MDB.DB()
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(global.Config.MySQLConfig.MaxOpenConns)
	db.SetMaxIdleConns(global.Config.MySQLConfig.MaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(global.Config.MySQLConfig.MaxLifetime) * time.Second)
}

func InitDB() error {

	baseDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=%s&parseTime=%t&loc=%s",
		global.Config.MySQLConfig.User,
		global.Config.MySQLConfig.Pass,
		global.Config.MySQLConfig.Host,
		strconv.Itoa(global.Config.MySQLConfig.Port),
		global.Config.MySQLConfig.Charset,
		global.Config.MySQLConfig.ParseTime,
		global.Config.MySQLConfig.Loc,
	)

	baseDB, err := gorm.Open(mysql.Open(baseDSN), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to connect to MySQL: %v", err))
	}

	dbName := global.Config.MySQLConfig.Name
	createDBSQL := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET %s", dbName, global.Config.MySQLConfig.Charset)
	if err := baseDB.Exec(createDBSQL).Error; err != nil {
		panic(fmt.Errorf("failed to create database: %v", err))
	}

	sqlDB, _ := baseDB.DB()
	sqlDB.Close()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		global.Config.MySQLConfig.User,
		global.Config.MySQLConfig.Pass,
		global.Config.MySQLConfig.Host,
		strconv.Itoa(global.Config.MySQLConfig.Port),
		dbName,
		global.Config.MySQLConfig.Charset,
		global.Config.MySQLConfig.ParseTime,
		global.Config.MySQLConfig.Loc,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to connect to database %s: %v", dbName, err))
	}

	global.MDB = db
	setPool()
	migrateTables()

	return nil
}
