// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"math"
// 	"math/rand"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// //Satellites
// type Satellites struct {
// 	SatellitesUnit []SatelliteUnit `json:"satellites"`
// }

// //SatelliteUnit
// type SatelliteUnit struct {
// 	Name     string   `json:"name"`
// 	Distance float64  `json:"distance"`
// 	Message  []string `json:"message"`
// }

// type ShipCoordinates struct {
// 	X string `json:"x"`
// 	Y string `json:"y"`
// }

// func main() {

// 	router := gin.Default()

// 	router.GET("/test", func(context *gin.Context) {
// 		context.JSON(200, gin.H{
// 			"MESSAGE": "ok!!!",
// 		})
// 	})

// 	router.POST("/topsecret_split/", func(context *gin.Context) {
// 		var satellite SatelliteUnit

// 		queryParam := context.Query("satellite_name")

// 		if strings.Count(queryParam, "") > 0 {
// 			requestBody := context.Request.Body
// 			x, _ := ioutil.ReadAll(requestBody)
// 			json.Unmarshal([]byte(x), &satellite)
// 			satellite.Name = queryParam

// 			fmt.Printf("\nTEST DO PARSER UCRANIANO- %#v\n", satellite)
// 			context.JSON(202, gin.H{

// 				"position": SpaceShipCoordinatesSattelitesUnit(satellite),
// 				"message":  MessageParserUnit(AsteroidInterferenceUnit(satellite)),
// 			})

// 		} else {
// 			context.JSON(404, gin.H{
// 				"status": "source not found",
// 			})
// 		}
// 	})

// 	router.POST("/topsecret", func(context *gin.Context) {

// 		var satellites Satellites

// 		requestBody := context.Request.Body
// 		x, _ := ioutil.ReadAll(requestBody)
// 		json.Unmarshal([]byte(x), &satellites)

// 		fmt.Printf("\nTEST DO PARSER UCRANIANO- %#v\n", satellites.SatellitesUnit[0])

// 		context.JSON(200, gin.H{
// 			"status":   "ok",
// 			"position": SpaceShipCoordinatesAllSattelites(satellites.SatellitesUnit),
// 			"message":  MessageParserAll(AsteroidInterference(satellites.SatellitesUnit)),
// 		})

// 	})

// 	router.Run(":8080") //listen and serve on 0.0.0.0:8080

// }

// func MessageParserAll(satellites []SatelliteUnit) string {

// 	messageParser := make([]string, len(satellites[0].Message))
// 	var i int
// 	for _, satelliteUnit := range satellites {
// 		// fmt.Printf("\n\n1o for: %#v", satelliteUnit)

// 		for _, message := range satelliteUnit.Message {
// 			// fmt.Printf("\n2o - for: valor de message: %#v\n", message)
// 			if message != "" {
// 				messageParser[i] = message

// 			}
// 			i++
// 		}
// 		i = 0

// 	}

// 	// fmt.Printf("\nArray Parseado: %#v", messageParser)
// 	return strings.Join(messageParser, " ")
// }

// func MessageParserUnit(satellite SatelliteUnit) string {

// 	messageParser := make([]string, len(satellite.Message))

// 	fmt.Printf("\n\n1o for: %#v", satellite)
// 	var i int
// 	for _, message := range satellite.Message {
// 		fmt.Printf("\n2o - for: valor de message: %#v\n", message)
// 		if message != "" && strings.ToLower(satellite.Name) == "kenobi" {
// 			messageParser[i] = message
// 		} else if message != "" && strings.ToLower(satellite.Name) == "skywalker" {
// 			messageParser[i] = message
// 		} else if message != "" && strings.ToLower(satellite.Name) == "sato" {
// 			messageParser[i] = message
// 		}
// 		i++

// 	}
// 	fmt.Printf("\nArray Parseado: %#v", messageParser)
// 	return strings.Join(messageParser, " ")

// }

// func AsteroidInterference(spaceShipTramission []SatelliteUnit) []SatelliteUnit {

// 	// fmt.Printf("\nSize do ramdom: %d", random)

// 	var i int
// 	for _, transmission := range spaceShipTramission {
// 		rand.Seed(time.Now().UnixNano())
// 		random := rand.Intn(len(spaceShipTramission[0].Message)-0+1) + 0
// 		for count := 0; count < random; count++ {
// 			transmission.Message[i] = ""
// 			// fmt.Printf("\n Dentro do Ramdom: %#v", transmission.Message[i])
// 			i++
// 		}
// 		i = 0
// 		//	fmt.Printf("\nArray Asteroid: %#v", spaceShip[i].Message)
// 	}

// 	// fmt.Printf("\nArray Asteroid: %#v\n", spaceShip[0].Message)
// 	return spaceShipTramission

// }

// func AsteroidInterferenceUnit(spaceShipTramission SatelliteUnit) SatelliteUnit {

// 	// fmt.Printf("\nSize do ramdom: %d", random)

// 	var i int
// 	for _, transmission := range spaceShipTramission.Message {
// 		rand.Seed(time.Now().UnixNano())
// 		random := rand.Intn(len(transmission)-0+1) + 0
// 		for count := 0; count < random; count++ {
// 			transmission = ""
// 			// fmt.Printf("\n Dentro do Ramdom: %#v", transmission.Message[i])
// 			i++
// 		}
// 		i = 0
// 		//	fmt.Printf("\nArray Asteroid: %#v", spaceShip[i].Message)
// 	}

// 	// fmt.Printf("\nArray Asteroid: %#v\n", spaceShip[0].Message)
// 	return spaceShipTramission

// }

// func SpaceShipCoordinatesAllSattelites(satellites []SatelliteUnit) ShipCoordinates {

// 	var shipCoordinates ShipCoordinates
// 	satoAngle := 30.00
// 	skywalkerAngle := 45.00
// 	kenobiAngle := 10.00

// 	type kenobiCoordinate struct {
// 		x float64
// 		y float64
// 	}

// 	type skywalkerCoordinate struct {
// 		x float64
// 		y float64
// 	}

// 	type satoCoordinate struct {
// 		x float64
// 		y float64
// 	}

// 	var kenobi kenobiCoordinate
// 	var skywalker skywalkerCoordinate
// 	var sato satoCoordinate

// 	var i int
// 	for _, satelliteUnit := range satellites {
// 		if strings.ToLower(satellites[i].Name) == "kenobi" {
// 			kenobi.x = math.Sin(kenobiAngle) * satelliteUnit.Distance
// 			kenobi.y = math.Cos(kenobiAngle) * satelliteUnit.Distance

// 			fmt.Printf("\nKenobi cordenada x: %.2f\n", kenobi.x)
// 			fmt.Printf("Kenobi cordenada y: %.2f\n", kenobi.y)

// 		} else if strings.ToLower(satellites[i].Name) == "skywalker" {
// 			skywalker.x = math.Sin(skywalkerAngle) * satelliteUnit.Distance
// 			skywalker.y = math.Cos(skywalkerAngle) * satelliteUnit.Distance

// 			fmt.Printf("\nskywalker cordenada x: %.2f\n", skywalker.x)
// 			fmt.Printf("skywalker cordenada y: %.2f\n", skywalker.y)

// 		} else if strings.ToLower(satellites[i].Name) == "sato" {
// 			sato.x = math.Sin(satoAngle) * satelliteUnit.Distance
// 			sato.y = math.Cos(satoAngle) * satelliteUnit.Distance

// 			fmt.Printf("\nsato cordenada x: %.2f\n", sato.x)
// 			fmt.Printf("sato cordenada y: %.2f\n", sato.y)

// 		}

// 		i++
// 	}

// 	shipCoordinates.X = strconv.FormatFloat(((kenobi.x + skywalker.x + sato.x) / 3), 'f', 2, 64)
// 	shipCoordinates.Y = strconv.FormatFloat(((kenobi.y + skywalker.y + sato.y) / 3), 'f', 2, 64)

// 	return shipCoordinates
// }

// func SpaceShipCoordinatesSattelitesUnit(satellite SatelliteUnit) ShipCoordinates {

// 	var shipCoordinates ShipCoordinates
// 	satoAngle := 30.00
// 	skywalkerAngle := 45.00
// 	kenobiAngle := 10.00

// 	if strings.ToLower(satellite.Name) == "kenobi" {
// 		kenobiXcoordinate := math.Sin(kenobiAngle) * satellite.Distance
// 		kenobiYcoordinate := math.Cos(kenobiAngle) * satellite.Distance
// 		shipCoordinates.X = strconv.FormatFloat((kenobiXcoordinate), 'f', 2, 64)
// 		shipCoordinates.Y = strconv.FormatFloat((kenobiYcoordinate), 'f', 2, 64)

// 		fmt.Printf("\nKenobi cordenada x: %.2f\n", kenobiXcoordinate)
// 		fmt.Printf("Kenobi cordenada y: %.2f\n", kenobiYcoordinate)

// 	} else if strings.ToLower(satellite.Name) == "skywalker" {
// 		skywalkerXcoordinate := math.Sin(skywalkerAngle) * satellite.Distance
// 		skywalkerYcoordinate := math.Cos(skywalkerAngle) * satellite.Distance
// 		shipCoordinates.X = strconv.FormatFloat((skywalkerXcoordinate), 'f', 2, 64)
// 		shipCoordinates.Y = strconv.FormatFloat((skywalkerYcoordinate), 'f', 2, 64)

// 		fmt.Printf("\nSkywalker cordenada x: %.2f\n", skywalkerXcoordinate)
// 		fmt.Printf("Skywalker cordenada y: %.2f\n", skywalkerYcoordinate)

// 	} else if strings.ToLower(satellite.Name) == "sato" {
// 		satoXcoordinate := math.Sin(satoAngle) * satellite.Distance
// 		satoYcoordinate := math.Cos(satoAngle) * satellite.Distance
// 		shipCoordinates.X = strconv.FormatFloat((satoXcoordinate), 'f', 2, 64)
// 		shipCoordinates.Y = strconv.FormatFloat((satoYcoordinate), 'f', 2, 64)

// 		fmt.Printf("\nSato cordenada x: %.2f\n", satoXcoordinate)
// 		fmt.Printf("Sato cordenada y: %.2f\n", satoYcoordinate)

// 	}

// 	fmt.Printf("Nave cordenada y: %#v\n", shipCoordinates)

// 	return shipCoordinates
// }
