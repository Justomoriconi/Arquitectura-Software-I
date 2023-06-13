package Controllers

import (
	Client "Backend/Clients/User"
	service "Backend/Services/Bookings"
	
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

func GetBookings(c *gin.Context) {
	log.Debug("loading bookings: ")

	hoteldomain, err := service.BookingService.GetBookings()

	if err != nil {
		c.JSON(http.StatusBadRequest, hoteldomain)
		return
	}

	c.JSON(http.StatusOK, hoteldomain)
}

/*
func GetmyBookings(c *gin.Context) {
	log.Debug("loading my bookings: ")
	user, _ := c.Get("user")
	id := user.(model.User).UserID

	bookingdomain, err := service.BookingService.GetmyBookings(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, bookingdomain)
		return
	}

	c.JSON(http.StatusOK, bookingdomain)
}*/

func GetmyBookings(c *gin.Context) {
	log.Debug("loading my bookings: ")
	id, _ := strconv.Atoi(c.Param("id"))

	bookingdomain, err := service.BookingService.GetmyBookings(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, bookingdomain)
		return
	}

	c.JSON(http.StatusOK, bookingdomain)
}

func Reserve(c *gin.Context) {
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
		fmt.Print(claims["exp"].(float64))
		//check expiration

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//find user

		User, err := Client.GetUserById(int(claims["sub"].(float64)))

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var Body struct {
			Hotelid  string
			Checkin  string
			Checkout string
		}
		if c.Bind(&Body) != nil {
			c.JSON(http.StatusBadRequest, Body)
			return
		}

		hotelid, _ := strconv.Atoi(Body.Hotelid)
		booking, err := service.BookingService.Reserve(User.UserID, hotelid, Body.Checkin, Body.Checkout)

		if err != nil {
			c.JSON(http.StatusBadRequest, booking)
			return
		}

		c.JSON(http.StatusOK, booking)
	}

	c.JSON(http.StatusBadRequest, gin.H{})
}

