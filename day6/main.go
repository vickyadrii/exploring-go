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
	ID           uint   `gorm:"primaryKey" json:"id"`
	Fullname     string `json:"fullname"`
	Username     string `json:"username"`
	PhotoProfile string `json:"photo_profile"`
}

type Tweets struct {
	ID     uint   `gorm:"primarykey" json:"id"`
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
	InitMigration()

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

	return c.JSON(http.StatusOK, map[string][]Users{
		"data": users,
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
	e.GET("/tweets", getTweets)

	e.Logger.Fatal(e.Start(":8000"))
	// fmt.Println("Day 6!!")
}
