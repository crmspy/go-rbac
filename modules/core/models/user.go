package models
 
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
//change table name
func (User) TableName() string {
    return "user"
}

type User struct {
	gorm.Model
	Username        string      `gorm:"unique" json:"username"`
	AuthKey         string      `json:"auth_key"`
	PasswordHash    string      `json:"password"`
	Email           string      `gorm:"unique" json:"email"`
	Status          string      `json:"status" sql:"DEFAULT:'active'"`
}
