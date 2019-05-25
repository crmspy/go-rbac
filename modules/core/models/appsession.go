package models
 
import (
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "time"
)
//change table name
func (ModelAppSession) TableName() string {
    return "app_session"
}

type ModelAppSession struct {
    AppSessionId    int		`gorm:"AUTO_INCREMENT;PRIMARY_KEY" json:"app_session_id"`
    AppKeyId        string      `gorm:"type:varchar(64);" `
    UserId          int         `json:"user_id"`
    UserAgent       string      `json:"user_agent"`
    Status          string      `gorm:"type:varchar(10);" json:"status" sql:"DEFAULT:'active'"`
    AccessToken     string      `json:"access_token"`
    LastAccess      time.Time   `json:"last_access"`
    CreatedBy       int         `json:"created_by"`
    CreatedAt       time.Time   `json:"created_at"`
    DeletedAt       time.Time   `json:"deleted_at"`
    ExpiredAt       time.Time   `json:"expired_at"`
}
