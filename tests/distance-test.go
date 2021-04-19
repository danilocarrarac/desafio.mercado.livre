// package main

// import (
// 	"fmt"
// 	"math"
// 	"strconv"
// )

// func main() {

// 	type ShipCoordinates struct {
// 		X string `json:"x"`
// 		Y string `json:"y"`
// 	}

// 	spaceShipDistance := 100.00
// 	var shipCoordinates ShipCoordinates
// 	satoAngle := 30.00
// 	skywalkerAngle := 45.00
// 	kenobiAngle := 10.00

// 	satoXcoordinate := math.Sin(satoAngle) * spaceShipDistance
// 	satoYcoordinate := math.Cos(satoAngle) * spaceShipDistance

// 	skywalkerXcoordinate := math.Sin(skywalkerAngle) * spaceShipDistance
// 	skywalkerYcoordinate := math.Cos(skywalkerAngle) * spaceShipDistance

// 	kenobiXcoordinate := math.Sin(kenobiAngle) * spaceShipDistance
// 	kenobiYcoordinate := math.Cos(kenobiAngle) * spaceShipDistance

// 	shipCoordinates.X = strconv.FormatFloat(((satoXcoordinate + skywalkerXcoordinate + kenobiXcoordinate) / 3), 'f', 2, 64)
// 	shipCoordinates.Y = strconv.FormatFloat(((satoYcoordinate + skywalkerYcoordinate + kenobiYcoordinate) / 3), 'f', 2, 64)

// 	fmt.Printf("\nSato cordenada x: %.2f\n", satoXcoordinate)
// 	fmt.Printf("Sato cordenada y: %.2f\n", satoYcoordinate)

// 	fmt.Printf("\nSkywalker cordenada x: %.2f\n", skywalkerXcoordinate)
// 	fmt.Printf("Skywalker cordenada y: %.2f\n", skywalkerYcoordinate)

// 	fmt.Printf("\nKenobi cordenada x: %.2f\n", kenobiXcoordinate)
// 	fmt.Printf("Kenobi cordenada y: %.2f\n", kenobiYcoordinate)

// 	fmt.Printf("Nave cordenada y: %#v\n", shipCoordinates)

// }

// // spaceShipdistance := 100.00

// // satoAngle := 30.00
// // skywalkerAngle := 45.00
// // kenobiAngle := 10.00

// // satoXcordinate := math.Sin(satoAngle) * spaceShipdistance
// // satoYcordinate := math.Cos(satoAngle) * spaceShipdistance

// // skywalkerXcordinate := math.Sin(skywalkerAngle) * spaceShipdistance
// // skywalkerYcordinate := math.Cos(skywalkerAngle) * spaceShipdistance

// // kenobiXcordinate := math.Sin(kenobiAngle) * spaceShipdistance
// // kenobiYcordinate := math.Cos(kenobiAngle) * spaceShipdistance

// // spaceShipXcordinate := (satoXcordinate + skywalkerXcordinate + kenobiYcordinate) / 3
// // spaceShipYcordinate := (satoYcordinate + skywalkerYcordinate + kenobiYcordinate) / 3

// // fmt.Printf("\nSato cordenada x: %.2f\n", satoXcordinate)
// // fmt.Printf("Sato cordenada y: %.2f\n", satoYcordinate)

// // fmt.Printf("\nSkywalker cordenada x: %.2f\n", skywalkerXcordinate)
// // fmt.Printf("Skywalker cordenada y: %.2f\n", skywalkerYcordinate)

// // fmt.Printf("\nKenobi cordenada x: %.2f\n", kenobiXcordinate)
// // fmt.Printf("Kenobi cordenada y: %.2f\n", kenobiYcordinate)

// // fmt.Printf("\nKenobi cordenada x: %.2f\n", kenobiXcordinate)
// // fmt.Printf("Kenobi cordenada y: %.2f\n", kenobiYcordinate)

// // fmt.Printf("\nNave cordenada x: %.2f\n", spaceShipXcordinate)
// // fmt.Printf("Nave cordenada y: %.2f\n", spaceShipYcordinate)
