package initialize

import (
	"errors"
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
		&persistent_object.ActivityLog{},
		&persistent_object.RefreshToken{},
	}
	err := global.MDB.AutoMigrate(models...)
	if err != nil {
		global.Logger.Error("Failed to migrate tables", zap.Error(err))
		panic(err)
	}
	seedRolesAndPermissions(global.MDB)
}

func seedRolesAndPermissions(db *gorm.DB) {
	permissions := []persistent_object.Permission{
		{Name: "video_upload", Description: "Allow to upload video"},
		{Name: "video_edit", Description: "Allow to edit video"},
		{Name: "video_delete", Description: "Allow to delete video"},
		{Name: "video_publish", Description: "Allow to publish video"},
		{Name: "channel_create", Description: "Allow to create channel"},
		{Name: "channel_edit", Description: "Allow to edit channel"},
		{Name: "channel_delete", Description: "Allow to delete channel"},
		{Name: "comment_post", Description: "Allow to post comment"},
		{Name: "comment_delete", Description: "Allow to delete comment"},
		{Name: "like_video", Description: "Allow to like video"},
		{Name: "report_content", Description: "Allow to report content"},
		{Name: "view_analytics", Description: "Allow to view analytics"},
		{Name: "manage_reports", Description: "Allow to manage reports"},
		{Name: "user_ban", Description: "Allow to ban user"},
		{Name: "user_edit", Description: "Allow to edit user"},
	}

	for _, perm := range permissions {
		var existing persistent_object.Permission
		if err := db.Where("name = ?", perm.Name).First(&existing).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if err := db.Create(&perm).Error; err != nil {
					global.Logger.Error("Failed to create permission", zap.String("permission", perm.Name), zap.Error(err))
				}
			} else {
				global.Logger.Error("Error querying permission", zap.String("permission", perm.Name), zap.Error(err))
			}
		}
	}

	var videoUpload, videoEdit, videoDelete, videoPublish persistent_object.Permission
	var channelCreate, channelEdit, channelDelete persistent_object.Permission
	var commentPost, commentDelete, likeVideo, reportContent persistent_object.Permission
	var viewAnalytics, manageReports, userBan, userEdit persistent_object.Permission

	db.Where("name = ?", "video_upload").First(&videoUpload)
	db.Where("name = ?", "video_edit").First(&videoEdit)
	db.Where("name = ?", "video_delete").First(&videoDelete)
	db.Where("name = ?", "video_publish").First(&videoPublish)
	db.Where("name = ?", "channel_create").First(&channelCreate)
	db.Where("name = ?", "channel_edit").First(&channelEdit)
	db.Where("name = ?", "channel_delete").First(&channelDelete)
	db.Where("name = ?", "comment_post").First(&commentPost)
	db.Where("name = ?", "comment_delete").First(&commentDelete)
	db.Where("name = ?", "like_video").First(&likeVideo)
	db.Where("name = ?", "report_content").First(&reportContent)
	db.Where("name = ?", "view_analytics").First(&viewAnalytics)
	db.Where("name = ?", "manage_reports").First(&manageReports)
	db.Where("name = ?", "user_ban").First(&userBan)
	db.Where("name = ?", "user_edit").First(&userEdit)

	// Định nghĩa các Role mặc định
	roles := []persistent_object.Role{
		{
			Name: "admin",
			Permissions: []*persistent_object.Permission{
				&videoUpload, &videoEdit, &videoDelete, &videoPublish,
				&channelCreate, &channelEdit, &channelDelete,
				&commentPost, &commentDelete, &likeVideo, &reportContent,
				&viewAnalytics, &manageReports, &userBan, &userEdit,
			},
		},
		{
			Name: "creator",
			Permissions: []*persistent_object.Permission{
				&videoUpload, &videoEdit, &videoPublish,
				&channelCreate, &channelEdit,
				&viewAnalytics,
				&commentPost, &likeVideo, &reportContent,
			},
		},
		{
			Name: "moderator",
			Permissions: []*persistent_object.Permission{
				&videoDelete, &commentDelete,
				&manageReports,
			},
		},
		{
			Name: "viewer",
			Permissions: []*persistent_object.Permission{
				&commentPost, &likeVideo, &reportContent,
			},
		},
	}

	for _, role := range roles {
		var existing persistent_object.Role
		if err := db.Where("name = ?", role.Name).First(&existing).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if err := db.Create(&role).Error; err != nil {
					global.Logger.Error("Failed to create role", zap.String("role", role.Name), zap.Error(err))
				}
			} else {
				global.Logger.Error("Error querying role", zap.String("role", role.Name), zap.Error(err))
			}
		}
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

	global.MDB = db
	setPool()
	migrateTables()

	return nil
}
