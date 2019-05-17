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
	
		if err == nil && token.Valid {
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
	Claims["expired"] = time.Now().Add(time.Hour * 72).Unix()

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
