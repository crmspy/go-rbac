package controller

import (
    ."github.com/crmspy/go-rbac/library/conn"
    core "github.com/crmspy/go-rbac/modules/core/models"
    cgx "github.com/crmspy/go-rbac/library/cgx"
    "errors"
    log "github.com/Sirupsen/logrus"
)

func CreateToken(appkeyid string,userid int,useragent string,createdby int,accesstoken string,expireddate int) (err error){

    appkey,err := GetAppKey(appkeyid)
    if (err != nil){
        return err
    }

    appsession := core.ModelAppSession{ AppKeyId:appkey.AppKeyId, UserId: 1, UserAgent:useragent, CreatedBy: createdby,CreatedAt: cgx.CgxNow(),AccessToken:accesstoken,ExpiredAt: cgx.CgxNowAdd(expireddate)}
    Db.Create(&appsession)
    return nil;
}

func AuthToken(token string) (err error){
    var datatoken core.ModelAppSession
    Db.Raw("select * from app_session where access_token = ?", token).Scan(&datatoken)
    if (datatoken.AccessToken==""){
        log.Warning("Token not valid")
        return errors.New("Token not valid")
    }else{
        if (datatoken.Status=="inactive"){
            return errors.New("Token inactive")
        }
        return nil
    }
}
