package controller

import (
    "github.com/gin-gonic/gin"
	"net/http"
    ."github.com/crmspy/go-rbac/library/conn"
    ."github.com/crmspy/go-rbac/modules/core/models"
    "errors"
    auth "github.com/crmspy/go-rbac/library/auth"
)
func GetProfile(c *gin.Context){
    var myToken string = c.GetHeader("Authorization");
    // parse token
    parsedToken,_ := auth.ParseToken(myToken)

    user := GetUsername(parsedToken["username"])
    c.JSON(http.StatusOK, gin.H{"status": "success", "data": user,"data_token":parsedToken})
}

func GetUsername(username interface{}) User{
    var user User
    Db.Where("username = ? ", username).First(&user)
    return user
   
}

func GetUser(userId int) User{
    var user User
    Db.First(&user, userId)
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
        }
        
        createdToken, err := auth.GenerateToken(payload)
        payload["token"]=createdToken
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
        }else{
            //return user with token
            c.JSON(http.StatusOK, gin.H{"status": "success", "data": payload})
        }
        }else{
        c.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
    }
}

func Auth(username string,password string) (User,error){
    var user User
    Db.Where("username = ? AND password_hash = ?", username, password).First(&user)
    if (user.Username==""){
        return user, errors.New("wrong username or password")
    }else{
        return user, nil
    }
}
