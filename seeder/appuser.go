package seeder

/* 
seeder function is generate a data into database
*/
import (
    cgx "github.com/crmspy/go-rbac/library/cgx"
    core "github.com/crmspy/go-rbac/modules/core/models"
    "github.com/crmspy/go-rbac/library/autoload"
    ."github.com/crmspy/go-rbac/library/conn"
)

func AppUser(){
    autoload.Run()
    //create new user as administrator
    password := cgx.CgxSha512("superadmin")
    authkey := cgx.CgxSha256(cgx.CgxGenerateString(64))
    var user = core.ModelAppUser{ Username: "superadmin", AuthKey: authkey, PasswordHash: password, Email: "crmspy@gmail.com", CreatedBy: 0, CreatedAt: cgx.CgxNow() }
    Db.Create(&user)
}
