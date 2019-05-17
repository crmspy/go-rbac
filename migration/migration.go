package migration
import (
    "github.com/crmspy/go-rbac/library/conn"
	log "github.com/Sirupsen/logrus"
    "github.com/fatih/structs"
    //inventory "github.com/crmspy/go-rbac/modules/inventory/models"
    //order "github.com/crmspy/go-rbac/modules/order/models"
    core "github.com/crmspy/go-rbac/modules/core/models"
)

/*
automatic migrate n seed data from model
*/

func Run(){
    //prepare create table and seed data
    // conn.Db.AutoMigrate(&inventory.Mproduct{})
    // conn.Db.AutoMigrate(&order.Torder{},&order.TorderLine{})
    // conn.Db.AutoMigrate(&inventory.Minventory{},&inventory.MinventoryLine{},&inventory.Tinout{})
    conn.Db.AutoMigrate(&core.ModelAppUser{})
    conn.Db.AutoMigrate(&core.ModelAppKey{})
    
}

func migrate(table interface{}){
    names := structs.Names(table)
    for _,t := range names {
        log.Println(t)
    }
}
