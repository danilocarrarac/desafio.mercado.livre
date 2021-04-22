package main

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"

	"go/go/src/github.com/danilocarrarac/desafio.mercado.livre/dtos"
	"go/go/src/github.com/danilocarrarac/desafio.mercado.livre/models"
	"go/go/src/github.com/danilocarrarac/desafio.mercado.livre/repository"
	"go/go/src/github.com/danilocarrarac/desafio.mercado.livre/services/asteroids"
	"go/go/src/github.com/danilocarrarac/desafio.mercado.livre/services/coordinateCalculation"
	"go/go/src/github.com/danilocarrarac/desafio.mercado.livre/services/messageParser"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.POST("/topsecret", func(context *gin.Context) {

		requestBody := new(models.Satellites)
		context.BindJSON(requestBody)

		context.JSON(200, gin.H{
			"status":   "ok",
			"position": coordinateCalculation.SpaceShipCoordinatesAllSattelites(requestBody.SatellitesUnit),
			"message":  messageParser.MessageParserAll(asteroids.AsteroidInterference(requestBody.SatellitesUnit)),
		})
	})

	router.GET("/topsecret_split/name/:satellite_name/distance/:distance", func(context *gin.Context) {

		if strings.Count(context.Param("satellite_name"), "") > 0 {
			var dbItem models.DBparser
			context.BindJSON(dbItem)
			dbItem.Name = context.Param("satellite_name")
			dbItem.Distance, _ = strconv.ParseFloat(context.Param("distance"), 64)

			dbConsult := repository.GetDBitem(dbItem)

			context.JSON(200, gin.H{

				"position": coordinateCalculation.SpaceShipCoordinatesSattelitesUnit((dtos.DTOdbParserToSatellite(dbConsult))),
				"message":  dbConsult.Message,
			})

		} else {
			context.JSON(404, gin.H{
				"status": "param should not be blank",
			})
		}

	})

	router.POST("/topsecret_split/name/:satellite_name", func(context *gin.Context) {
		var satelliteUnit models.SatelliteUnit
		var dbInsert models.DBparser

		if strings.Count(context.Param("satellite_name"), "") > 0 {
			requestBody := context.Request.Body
			x, _ := ioutil.ReadAll(requestBody)
			json.Unmarshal([]byte(x), &satelliteUnit)

			satelliteUnit.Name = context.Param("satellite_name")

			context.JSON(200, gin.H{

				"position": coordinateCalculation.SpaceShipCoordinatesSattelitesUnit(satelliteUnit),
				"message":  messageParser.MessageParserUnit(asteroids.AsteroidInterferenceUnit(satelliteUnit)),
			})

			dbInsert = dtos.DTOsatelliteToDBparser(satelliteUnit)
			repository.PostDBitem(dbInsert)

		} else {
			context.JSON(404, gin.H{
				"status": "source not found",
			})
		}
	})

	router.Run(":3000")

}
