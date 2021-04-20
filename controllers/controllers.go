package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"go/go/src/github.com/danilocarrarac/desafio.mercado.livre/models"
	"go/go/src/github.com/danilocarrarac/desafio.mercado.livre/repository"
	"go/go/src/github.com/danilocarrarac/desafio.mercado.livre/services/asteroids"
	"go/go/src/github.com/danilocarrarac/desafio.mercado.livre/services/coordinateCalculation"
	"go/go/src/github.com/danilocarrarac/desafio.mercado.livre/services/messageParser"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/topsecret_split/", func(context *gin.Context) {

		var dbConsult models.DBconsult
		var satelliteUnit models.SatelliteUnit
		sateliteName := context.Query("satellite_name")
		distance := context.Query("distance")

		if strings.Count(sateliteName, "") > 0 {
			requestBody := context.Request.Body
			x, _ := ioutil.ReadAll(requestBody)
			json.Unmarshal([]byte(x), &dbConsult)
			dbConsult.Name = sateliteName
			dbConsult.Distance, _ = strconv.ParseFloat(distance, 64)

			fmt.Printf("\nvalor de satelite no get: %#v: \n", dbConsult)

			dbItem := repository.GetDBitem(dbConsult)

			satelliteUnit.Name = dbItem.Name
			satelliteUnit.Distance = dbItem.Distance

			context.JSON(202, gin.H{

				"position": coordinateCalculation.SpaceShipCoordinatesSattelitesUnit((satelliteUnit)),
				"message":  dbItem.Message,
			})

		} else {
			context.JSON(404, gin.H{
				"status": "source not found",
			})
		}

	})

	router.POST("/topsecret_split/", func(context *gin.Context) {
		var satellite models.SatelliteUnit
		var dbInsert models.DBconsult

		queryParam := context.Query("satellite_name")

		if strings.Count(queryParam, "") > 0 {
			requestBody := context.Request.Body
			x, _ := ioutil.ReadAll(requestBody)
			json.Unmarshal([]byte(x), &satellite)
			satellite.Name = queryParam

			context.JSON(202, gin.H{

				"position": coordinateCalculation.SpaceShipCoordinatesSattelitesUnit(satellite),
				"message":  messageParser.MessageParserUnit(asteroids.AsteroidInterferenceUnit(satellite)),
			})

			dbInsert.Name = satellite.Name
			dbInsert.Message = messageParser.MessageParserUnit(asteroids.AsteroidInterferenceUnit(satellite))
			dbInsert.Distance = satellite.Distance

			repository.PostDBitem(dbInsert)

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
