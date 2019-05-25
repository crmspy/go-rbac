package controller

import (
    ."github.com/crmspy/go-rbac/library/conn"
    core "github.com/crmspy/go-rbac/modules/core/models"
    "errors"
    log "github.com/Sirupsen/logrus"
)

func GetAppKey(appkeyid string) (core.ModelAppKey,error){
    var appkey core.ModelAppKey
    Db.Raw("select * from app_key where app_key_id = ?", appkeyid).Scan(&appkey)
    if (appkey.AppKeyId==""){
        log.Warning("appkey not found'")
        return appkey, errors.New("appkey not found")
    }else{
        return appkey, nil
    }
}
