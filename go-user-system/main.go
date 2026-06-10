package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

var db *gorm.DB
var jwtSecret = []byte("simple-super-secret-key")

func main() {
	database()
	router := gin.Default()

	user := router.Group("/user")
	user.Use(AuthMiddleware())
	{
		user.GET("/:id", getUser)
		user.POST("/add", addUser)
		user.PUT("/:id", updateUser)
		user.DELETE("/:id", removeUser)
		//		user.POST("/token", )
	}

	router.POST("/authenticate", authenticate)

	err := router.Run(":8081")

	if err != nil {
		fmt.Println("error running server:", err)
		panic(err)
	}
}

type User struct {
	UserId    uint   `gorm:"column:user_id;primaryKey;autoIncrement" json:"user_id"`
	FirstName string `gorm:"column:firstname;" json:"first_name"`
	LastName  string `gorm:"column:lastname" json:"last_name"`
	Username  string `gorm:"column:username" json:"username"`
	Email     string `gorm:"column:email" json:"email"`
	Password  string `gorm:"column:password" json:"password"`
}

func database() {
	var err error
	var dns = "host=127.0.0.1 user=root password=root123 dbname=go_user port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dns))
	if err != nil {
		fmt.Println("failed to connect database")
	}
}

func getUser(c *gin.Context) {
	var userId = c.Param("id")

	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	var user User
	res := db.First(&user, "user_id = ?", userId)

	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func addUser(c *gin.Context) {

	var user User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record := User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
	}

	if len(user.Password) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password must be at least 8 characters"})
		return
	}

	if len(user.Email) < 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email must be at least 5 characters"})
		return
	}

	pass, err := bcryptPassword(user.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record.Password = pass

	var res = db.Create(&record)

	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": record})
}

func updateUser(c *gin.Context) {
	var userId = c.Param("id")

	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	var user User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	record := User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
	}

	if len(user.Password) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password must be at least 8 characters"})
		return
	}

	if len(user.Email) < 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email must be at least 5 characters"})
		return
	}

	pass, err := bcryptPassword(user.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record.Password = pass

	var res = db.Model(&User{}).Where("user_id = ?", userId).Updates(&record)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": record})
}

func removeUser(c *gin.Context) {
	var userId = c.Param("id")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	var user User
	res := db.Where("user_id = ?", userId).Delete(&user)

	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_id": userId})
}

type AuthenticateRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func authenticate(c *gin.Context) {
	var authenticateRequest AuthenticateRequest

	if err := c.ShouldBindJSON(&authenticateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var password = authenticateRequest.Password
	var username = authenticateRequest.Username

	var user User
	var res = db.First(&user, "username = ?", username)

	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
	}

	if !validatePassword(user.Password, password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}

	token, err := jwtToken(user.UserId, user.Username, true)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user, "token": token})

}

func bcryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(bytes), err
}

func validatePassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func jwtToken(userID uint, username string, allowed bool) (string, error) {

	return "", nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
