package in_memory_db

import (
	"errors"
	"fmt"

	"github.com/bhtoan2204/user/global"
	"github.com/bhtoan2204/user/internal/infrastructure/db/in_memory_db/persistent_object_test"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func seedRolesAndPermissions(db *gorm.DB) {
	permissions := []persistent_object_test.Permission{
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
		var existing persistent_object_test.Permission
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

	var videoUpload, videoEdit, videoDelete, videoPublish persistent_object_test.Permission
	var channelCreate, channelEdit, channelDelete persistent_object_test.Permission
	var commentPost, commentDelete, likeVideo, reportContent persistent_object_test.Permission
	var viewAnalytics, manageReports, userBan, userEdit persistent_object_test.Permission

	// db.Where("name = ?", "video_upload").First(&videoUpload)
	// db.Where("name = ?", "video_edit").First(&videoEdit)
	// db.Where("name = ?", "video_delete").First(&videoDelete)
	// db.Where("name = ?", "video_publish").First(&videoPublish)
	// db.Where("name = ?", "channel_create").First(&channelCreate)
	// db.Where("name = ?", "channel_edit").First(&channelEdit)
	// db.Where("name = ?", "channel_delete").First(&channelDelete)
	// db.Where("name = ?", "comment_post").First(&commentPost)
	// db.Where("name = ?", "comment_delete").First(&commentDelete)
	// db.Where("name = ?", "like_video").First(&likeVideo)
	// db.Where("name = ?", "report_content").First(&reportContent)
	// db.Where("name = ?", "view_analytics").First(&viewAnalytics)
	// db.Where("name = ?", "manage_reports").First(&manageReports)
	// db.Where("name = ?", "user_ban").First(&userBan)
	// db.Where("name = ?", "user_edit").First(&userEdit)

	roles := []persistent_object_test.Role{
		{
			Name: "admin",
			Permissions: []*persistent_object_test.Permission{
				&videoUpload, &videoEdit, &videoDelete, &videoPublish,
				&channelCreate, &channelEdit, &channelDelete,
				&commentPost, &commentDelete, &likeVideo, &reportContent,
				&viewAnalytics, &manageReports, &userBan, &userEdit,
			},
		},
		{
			Name: "creator",
			Permissions: []*persistent_object_test.Permission{
				&videoUpload, &videoEdit, &videoPublish,
				&channelCreate, &channelEdit,
				&viewAnalytics,
				&commentPost, &likeVideo, &reportContent,
			},
		},
		{
			Name: "moderator",
			Permissions: []*persistent_object_test.Permission{
				&videoDelete, &commentDelete,
				&manageReports,
			},
		},
		{
			Name: "viewer",
			Permissions: []*persistent_object_test.Permission{
				&commentPost, &likeVideo, &reportContent,
			},
		},
	}

	for _, role := range roles {
		var existing persistent_object_test.Role
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

func CreateTestDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the test database")
	}
	models := []interface{}{
		&persistent_object_test.User{},
		&persistent_object_test.Role{},
		&persistent_object_test.ActivityLog{},
		&persistent_object_test.Permission{},
		&persistent_object_test.UserSettings{},
		&persistent_object_test.RefreshToken{},
	}

	// Auto migrate tables
	err = db.AutoMigrate(models...)
	if err != nil {
		fmt.Printf("Failed to migrate tables: %v", err)
		panic(err)
	}

	// Verify tables are created
	for _, model := range models {
		if !db.Migrator().HasTable(model) {
			panic(fmt.Sprintf("Table for %T was not created", model))
		}
	}

	// Seed data only after tables are confirmed to exist
	seedRolesAndPermissions(db)
	return db
}
