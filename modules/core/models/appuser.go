package models
 
import (
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "time"
)
//change table name
func (ModelAppUser) TableName() string {
    return "app_user"
}

type ModelAppUser struct {
    AppUserId       int         `gorm:"AUTO_INCREMENT;PRIMARY_KEY;COLUMN:app_user_id"  json:"user_id"`
    Username        string      `gorm:"unique;type:varchar(32);" json:"username"`
	AuthKey         string      `json:"auth_key"`
	PasswordHash    string      `json:"password"`
	Email           string      `gorm:"unique;type:varchar(64);" json:"email"`
    Status          string      `gorm:"type:varchar(10);" json:"status" sql:"DEFAULT:'active'"`
    CreatedBy       int         `json:"created_by"`
    UpdatedBy       int         `json:"updated_by"`
    CreatedAt       time.Time   `json:"created_at"`
    UpdatedAt       time.Time   `json:"updated_at"`
    DeletedAt       time.Time   `json:"deleted_at"`
}
