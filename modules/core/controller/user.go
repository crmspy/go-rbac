package controller

import (
    "github.com/gin-gonic/gin"
	"net/http"
    ."github.com/crmspy/go-rbac/library/conn"
    core "github.com/crmspy/go-rbac/modules/core/models"
    "errors"
    auth "github.com/crmspy/go-rbac/library/auth"
    cgx "github.com/crmspy/go-rbac/library/cgx"
    "github.com/spf13/viper"
    "strconv"
)

func GetProfile(c *gin.Context){
    var myToken string = c.GetHeader("Authorization");
    // parse token
    parsedToken,_ := auth.ParseToken(myToken)

    user := GetUsername(parsedToken["username"])
    c.JSON(http.StatusOK, gin.H{"status": "success", "data": user,"data_token":parsedToken})
}

func GetUsername(username interface{}) core.ModelAppUser{
    var user core.ModelAppUser
    Db.Raw("SELECT * FROM app_user WHERE username = ?", username).Scan(&user)
    return user
   
}

func GetUser(userId int) core.ModelAppUser{
    var user core.ModelAppUser
    Db.Raw("SELECT * FROM app_user WHERE app_user_id = ?", userId).Scan(&user)
    return user
}

func Login(c *gin.Context){
    username := c.PostForm("username")
    password := c.PostForm("password")
    user,err := Auth(username,password)
    
    if (err == nil){
        var payload = map[string]string{
            "username": user.Username,
            "email": user.Email,
            "user_id": strconv.Itoa(user.AppUserId),
        }
        
        createdToken, err := auth.GenerateToken(payload)

        var ret = map[string]string{
            "token": createdToken,
        }
        
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
        }
        
        //create new session and add to db
        expireddate, err := strconv.Atoi(viper.GetString("session_expired"))
        useragent := c.Request.UserAgent()
        err = CreateToken("WEB",user.AppUserId,useragent,user.AppUserId,createdToken,expireddate)

        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
        }else{
            //return user with token
            c.JSON(http.StatusOK, gin.H{"status": "success", "data": ret})
        }
    }else{
        c.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
    }
}

func Auth(username string,password string) (core.ModelAppUser,error){
    var user core.ModelAppUser
    var hash_password string = cgx.CgxSha512(password) 

    Db.Raw("SELECT * FROM app_user WHERE username = ? and password_hash = ?",username, hash_password).Scan(&user)
    if (user.Username==""){
        return user, errors.New("wrong username or password")
    }else{
        return user, nil
    }
}

func Test(c *gin.Context){
    // err := CreateSession();
    // if (err != nil){
    //     c.JSON(http.StatusOK, gin.H{"status": "fail", "data": ""})
    // }else{
    //     c.JSON(http.StatusOK, gin.H{"status": "success", "data": "1"})
    // }
    
}
