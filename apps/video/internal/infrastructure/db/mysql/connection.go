package mysql

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/bhtoan2204/video/global"
	"github.com/bhtoan2204/video/internal/infrastructure/db/mysql/persistent_object"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func migrateTables() {
	models := []interface{}{
		&persistent_object.PlaylistVideo{},
		&persistent_object.Playlist{},
		&persistent_object.UserVideoMetadata{},
		&persistent_object.VideoMetadata{},
		&persistent_object.Tag{},
		&persistent_object.Video{},
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
		global.Logger.Error("Failed to connect to MySQL", zap.Error(err))
		panic(err)
	}

	dbName := global.Config.MySQLConfig.Name
	createDBSQL := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET %s", dbName, global.Config.MySQLConfig.Charset)
	if err := baseDB.Exec(createDBSQL).Error; err != nil {
		global.Logger.Error("Failed to create database", zap.Error(err))
		panic(err)
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
		global.Logger.Error("Failed to connect to MySQL", zap.Error(err))
		panic(err)
	}

	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		log.Fatalf("Failed to use otelgorm plugin: %v", err)
	}

	global.MDB = db
	setPool()
	migrateTables()

	return nil
}

func CloseDB() {
	if global.MDB != nil {
		sqlDB, err := global.MDB.DB()
		if err != nil {
			global.Logger.Error("Failed to get database connection", zap.Error(err))
			return
		}
		if err := sqlDB.Close(); err != nil {
			global.Logger.Error("Failed to close database connection", zap.Error(err))
		} else {
			global.Logger.Info("Database connection closed")
		}
	}
}
