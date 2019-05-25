package models
 
import (
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "time"
)
//change table name
func (ModelAppKey) TableName() string {
    return "app_key"
}

type ModelAppKey struct {
    AppKeyId	    string		`gorm:"type:varchar(64);PRIMARY_KEY" json:"key_id"`
    Key             string      `gorm:"unique;type:varchar(32);" json:"email"`
    Name            string      `gorm:"type:varchar(64);" json:"name"`
    Status          string      `gorm:"type:varchar(10);" json:"status" sql:"DEFAULT:'active'"`
    GeneratedKey    string      `json:"generated_key"`
    LastAccess      time.Time   `json:"last_access"`
    CreatedBy      int         `json:"created_by"`
    UpdatedBy      int         `json:"updated_by"`
    CreatedAt       time.Time   `json:"created_at"`
    UpdatedAt       time.Time   `json:"updated_at"`
    DeletedAt       time.Time   `json:"deleted_at"`
}
