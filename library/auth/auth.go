package auth

/*
Code By Nurul Hidayat
crmspy@gmail.com

authentication using jwt with database support
*/

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
    "github.com/dgrijalva/jwt-go"
    "github.com/spf13/viper"
    "errors"
    core "github.com/crmspy/go-rbac/modules/core/models"
    ."github.com/crmspy/go-rbac/library/conn"
    log "github.com/Sirupsen/logrus"
    cgx "github.com/crmspy/go-rbac/library/cgx"
    "strconv"
)

func AuthError(code int, message string,c *gin.Context ) {
	resp := map[string]string{"error": message}
	c.JSON(code, resp)
	c.Abort()
  }

func Auth(c *gin.Context) {
	var myToken string = c.GetHeader("Authorization");
	if myToken != "" {
		token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("generate_key")), nil
		})
        
        tokenvalid := SessionCheck(myToken);
        if tokenvalid != nil{
            AuthError(401,tokenvalid.Error(),c)
        }else if err == nil && token.Valid {
			c.Next()
		}else{
		    AuthError(401,"wrong token",c)
		}
	}else{
		    AuthError(401,"wrong token",c)
	}

}



func GetKey(c *gin.Context) {
    createdToken, err := SampleGenerateToken([]byte(viper.GetString("generate_key")))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "Creating TOken failed",
		})
    }else{
		c.JSON(200, gin.H{
			"message": createdToken,
		})
	}
 
}

func GenerateToken(payload map[string]string) (string, error) {
    mySigningKey := []byte(viper.GetString("generate_key"))
    // Create the token
    token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
    Claims := make(jwt.MapClaims)
    for key,val := range payload{
        Claims[key] = val
    }
    
    //set expired time
    expiredminute,_ := strconv.Atoi(viper.GetString("session_expired"))
	Claims["expired"] = cgx.CgxNowAdd(expiredminute)

	token.Claims = Claims
    // Sign and get the complete encoded token as a string
    tokenString, err := token.SignedString(mySigningKey)
    return tokenString, err
}



func ParseToken(myToken string) (jwt.MapClaims,error){
    token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
        return []byte(viper.GetString("generate_key")), nil
    })

    if err == nil && token.Valid {
        user := token.Claims.(jwt.MapClaims)
        return user,err
    } else {
        return nil,errors.New("This token is terrible!  I cannot accept this.")
    }
}

func SampleGenerateToken(mySigningKey []byte) (string, error) {
    // Create the token
    token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	Claims := make(jwt.MapClaims)
	Claims["foo"] = "Okey Im Coming"
	Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token.Claims = Claims
    // Sign and get the complete encoded token as a string
    tokenString, err := token.SignedString(mySigningKey)
    return tokenString, err
}

func SessionCheck(token string) (err error){
    var datatoken core.ModelAppSession
    var datatokenexp core.ModelAppSession
    Db.Raw("select * from app_session where access_token = ?", token).Scan(&datatoken)
    Db.Raw("select * from app_session where access_token = ? and (UNIX_TIMESTAMP(expired_at) = 0 or expired_at > now())", token).Scan(&datatokenexp)
    
    if (datatoken.AccessToken==""){
        log.Warning("Token not valid")
        return errors.New("Token not valid")
    }else{
        if (datatoken.Status=="inactive"){
            return errors.New("Token inactive")
        }else if (datatoken.ExpiredAt == time.Time{}){
            Db.Exec("UPDATE app_session SET last_access=? WHERE access_token = ?",cgx.CgxNow(), token)
            return nil
        }else if (datatokenexp.AccessToken==""){
            Db.Exec("UPDATE app_session SET status='inactive',deleted_at=? WHERE access_token = ?",cgx.CgxNow(), token)
            return errors.New("Token expired")
        }
        return nil
    }
}
