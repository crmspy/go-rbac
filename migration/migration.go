package migration
import (
    "github.com/crmspy/go-rbac/library/conn"
	log "github.com/Sirupsen/logrus"
    "github.com/fatih/structs"
    core "github.com/crmspy/go-rbac/modules/core/models"
)

/*
automatic migrate n seed data from model
*/

func Run(){
    conn.Db.AutoMigrate(&core.ModelAppUser{})
    conn.Db.AutoMigrate(&core.ModelAppKey{})
    conn.Db.AutoMigrate(&core.ModelAppSession{})
    
}

func migrate(table interface{}){
    names := structs.Names(table)
    for _,t := range names {
        log.Println(t)
    }
}
