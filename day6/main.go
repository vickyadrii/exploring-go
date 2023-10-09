package main

import (
	// "fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Users struct {
	// gorm.Model
	ID           uint   `gorm:"primaryKey" json:"id" form:"id"`
	Fullname     string `json:"fullname" form:"fullname"`
	Username     string `json:"username" form:"username"`
	PhotoProfile string `json:"photo_profile" form:"photo_profile"`
}

type Tweets struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Tweet  string `json:"tweet"`
	UserId uint   `json:"user_id"`
}

func InitDatabase() {
	dsn := "root:@tcp(127.0.0.1:3306)/prakerja12?charset=utf8mb4&parseTime=True&loc=Local" // Database Source Name
	var err error
	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

}

func InitMigration() {
	DB.AutoMigrate()
}

func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}

func getUsersController(c echo.Context) error {
	var users []Users
	result := DB.Find(&users)

	if result.Error != nil {
		return result.Error
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"data":    users,
		"message": "success get all users",
	})
}

func CreateUserController(c echo.Context) error {
	users := Users{}
	c.Bind(&users)

	payload := DB.Save(&users)

	if payload.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, payload.Error)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"data":    users,
	})
}

func getTweets(c echo.Context) error {
	var tweets []Tweets
	result := DB.Find(&tweets)

	if result.Error != nil {
		return result.Error
	}

	return c.JSON(http.StatusOK, map[string][]Tweets{
		"data": tweets,
	})
}

func main() {
	InitDatabase()
	e := echo.New() // Create instance
	e.GET("/", Home)
	e.GET("/users", getUsersController)
	e.POST("/users", CreateUserController)
	e.GET("/tweets", getTweets)

	e.Logger.Fatal(e.Start(":8000"))
	// fmt.Println("Day 6!!")
}
