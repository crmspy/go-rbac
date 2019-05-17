package seeder

import (
    cgx "github.com/crmspy/go-rbac/library/cgx"
    core "github.com/crmspy/go-rbac/modules/core/models"
    ."github.com/crmspy/go-rbac/library/conn"
    "github.com/crmspy/go-rbac/library/autoload"
)
func AppUser(){
    autoload.Run()
    password := cgx.CgxSha512("superadmin")
    authkey := cgx.CgxSha256("crmspy@gmail.com")
    var user = core.ModelAppUser{ Username: "superadmin", AuthKey: authkey, PasswordHash: password, Email: "crmspy@gmail.com", CreatedBy: 0, CreatedAt: cgx.CgxNow() }
    Db.Create(&user)
}
