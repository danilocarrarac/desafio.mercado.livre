package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"desafio.mercado.livre/services/asteroids"
	"desafio.mercado.livre/services/coordinateCalculation"
	"desafio.mercado.livre/services/messageParser"
	"github.com/danilocarrarac/desafio.mercado.livre/models"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/test", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"MESSAGE": "ok!!!",
		})
	})

	router.POST("/topsecret_split/", func(context *gin.Context) {
		var satellite models.SatelliteUnit

		queryParam := context.Query("satellite_name")

		if strings.Count(queryParam, "") > 0 {
			requestBody := context.Request.Body
			x, _ := ioutil.ReadAll(requestBody)
			json.Unmarshal([]byte(x), &satellite)
			satellite.Name = queryParam

			fmt.Printf("\nTEST DO PARSER UCRANIANO- %#v\n", satellite)
			context.JSON(202, gin.H{

				"position": coordinateCalculation.SpaceShipCoordinatesSattelitesUnit(satellite),
				"message":  messageParser.MessageParserUnit(asteroids.AsteroidInterferenceUnit(satellite)),
			})

		} else {
			context.JSON(404, gin.H{
				"status": "source not found",
			})
		}
	})

	router.POST("/topsecret", func(context *gin.Context) {

		var satellites models.Satellites

		requestBody := context.Request.Body
		x, _ := ioutil.ReadAll(requestBody)
		json.Unmarshal([]byte(x), &satellites)

		context.JSON(200, gin.H{
			"status":   "ok",
			"position": coordinateCalculation.SpaceShipCoordinatesAllSattelites(satellites.SatellitesUnit),
			"message":  messageParser.MessageParserAll(asteroids.AsteroidInterference(satellites.SatellitesUnit)),
		})

	})

	router.Run(":8080") //listen and serve on 0.0.0.0:8080

}
