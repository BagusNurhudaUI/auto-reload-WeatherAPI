package main

import (
	"fmt"
	"html/template"
	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/jasonlvhit/gocron"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

var status Status

func main() {

	status = Status{rand.Intn(100), rand.Intn(100)}
	gocron.Every(15).Seconds().Do(getUpdatedData)
	gocron.Start()
	fmt.Println("Welcome to the server!")
	r := gin.Default()
	r.GET("/", indexPage)
	r.GET("/status", StatusWeather)
	r.Run("127.0.0.1:3000")
}

func indexPage(c *gin.Context) {
	result := gin.H{
		"status": status,
	}
	c.JSON(200, result)
}

func getUpdatedData() {
	status = Status{rand.Intn(100), rand.Intn(100)}
	fmt.Println("I am running to update : Water, Wind :", status)
}

func StatusWeather(c *gin.Context) {
	tmpl := template.Must(template.ParseFiles("main.html"))
	tmpl.Execute(c.Writer, status)
}
