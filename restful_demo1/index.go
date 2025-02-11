package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ----------------------------------------------
type User struct {
	Username string `jason: "Username"`
	Password string `jason: "Password"`
}

// ----------------------------------------------
func GetHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
}

func GetUser(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	c.JSON(http.StatusOK, gin.H{"username": username, "password": password})
}

// ----------------------------------------------
func PostUser(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User": user})
}

// ----------------------------------------------
func PutPathParams(c *gin.Context) {
	user := c.Param("user")
	c.JSON(http.StatusOK, gin.H{"message": user})
}

// ----------------------------------------------
func main() {
	r := gin.Default()
	r.GET("/", GetHello)
	r.GET("/user", GetUser)
	r.POST("/", PostUser)
	r.Run()
}
