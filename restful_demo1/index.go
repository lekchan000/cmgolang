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
	username := c.Param("username")
	password := c.Param("password")
	c.JSON(http.StatusOK, gin.H{"username": username, "password": password})
}

// ----------------------------------------------
func DeleteUser(c *gin.Context) {
	username := c.Query("username")
	c.JSON(http.StatusOK, gin.H{"username": username})
}

// ----------------------------------------------
func main() {
	r := gin.Default()
	r.GET("/", GetHello)
	r.GET("/user", GetUser)
	r.POST("/", PostUser)
	r.PUT("/user/:username/:password", PutPathParams)
	r.DELETE("/username", DeleteUser)
	r.Run()
}
