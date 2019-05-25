package controller

import (
    ."github.com/crmspy/go-rbac/library/conn"
    core "github.com/crmspy/go-rbac/modules/core/models"
    cgx "github.com/crmspy/go-rbac/library/cgx"

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
