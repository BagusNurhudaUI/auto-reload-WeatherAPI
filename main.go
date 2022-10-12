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

var (
	angin  string
	air    string
	status Status
)

func main() {
	fmt.Println("Welcome to the server!")
	getUpdatedData() //get first data
	gocron.Every(15).Seconds().Do(getUpdatedData)
	gocron.Start()

	r := gin.Default()
	r.GET("/", StatusWeather)
	r.GET("/status", indexPage)
	r.Run("127.0.0.1:3000")
}

func indexPage(c *gin.Context) {
	result := gin.H{
		"status": status,
	}
	c.JSON(200, result)
}

func getStatus() {
	status = Status{rand.Intn(100), rand.Intn(100)}
	if status.Wind <= 5 {
		angin = "Aman"
	} else if status.Wind < 8 {
		angin = "Siaga"
	} else {
		angin = "Bahaya"
	}

	if status.Water <= 6 {
		air = "Aman"
	} else if status.Water < 15 {
		air = "Siaga"
	} else {
		air = "Bahaya"
	}
}
func getUpdatedData() {
	getStatus()
	fmt.Println("--------------------------------")
	fmt.Println("Running the update....")
	fmt.Println("Wind:", status.Wind, "  Status:", angin)
	fmt.Println("Water:", status.Water, "  Status:", air)

}

func StatusWeather(c *gin.Context) {
	tmpl := template.Must(template.ParseFiles("main.html"))
	tmpl.Execute(c.Writer, status)
}
