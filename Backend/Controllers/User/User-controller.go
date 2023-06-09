package Controllers

import (
	Client "Backend/Clients/User"
	service "Backend/Services/User"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
)

// Gets
func GetUserById(c *gin.Context) {
	log.Debug("User id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	userdomain, err := service.UserService.GetUserById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, userdomain)
		return
	}

	c.JSON(http.StatusOK, userdomain)
}

func GetUSerByusername(c *gin.Context) {
	log.Debug("User name to load: " + c.Param("username"))

	username := c.Param("username")
	userdomain, err := service.UserService.GetUSerByusername(username)

	if err != nil {
		c.JSON(http.StatusBadRequest, userdomain)
		return
	}

	c.JSON(http.StatusOK, userdomain)
}

// Singup
func Singup(c *gin.Context) {
	//get singup data
	var Body struct {
		Name     string
		LastName string
		Email    string
		Pwd      string
	}

	if c.Bind(&Body) != nil {
		c.JSON(http.StatusBadRequest, Body)
		return
	}
	log.Debug()
	//hash pasword
	hash, err := bcrypt.GenerateFromPassword([]byte(Body.Pwd), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, Body)
		return
	}
	//create user
	err = service.UserService.Singup(Body.Name, Body.LastName, Body.Email, string(hash))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{}) //empty return
}

func Login(c *gin.Context) {
	//get pwd and email
	var Body struct {
		Email string
		Pwd   string
	}

	if c.Bind(&Body) != nil {
		c.JSON(http.StatusBadRequest, Body)
		return
	}

	//look for user
	userdomain, err := service.UserService.GetUSerByEmail(Body.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "1invalid email or password"})
		return
	}
	//compare pwd
	err = bcrypt.CompareHashAndPassword([]byte(userdomain.Pwd), []byte(Body.Pwd))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "2invalid email or password"})
		return
	}
	//genereta token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userdomain.UserID,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte("ajfgnaigingeiuaw"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "cannot generate token"})
		return
	}
	//send it
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"mesage": "im loged in"})
}

func Logout(c *gin.Context) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{

		"exp": time.Now(),
	})

	tokenString, err := token.SignedString([]byte("ajfgnaigingeiuaw"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "cannot generate token"})
		return
	}
	//send it
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 0, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})

}
func Myuser(c *gin.Context) {

	//Get cookie
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	//validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("ajfgnaigingeiuaw"), nil
	})
	if err != nil || token == nil || !token.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		//check expiration

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//find user

		User, err := Client.GetUserById(int(claims["sub"].(float64)))

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		//atach to request
		c.JSON(http.StatusOK, gin.H{"userid": User.UserID})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{})
	}
}
