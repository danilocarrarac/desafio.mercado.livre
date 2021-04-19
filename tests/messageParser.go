// package main

// import (
// 	"fmt"
// )

// // s := new(Satellites)
// func messageParser(ss []string) {

// 	fmt.Printf("%#v: ", ss[0])
// }

// func main() {

// 	kenobi := []string{"", "es", "un", "mensaje"}
// 	skywalker := []string{"este", "", "un", "mensaje"}
// 	sato := []string{"", "es", "", "mensaje"}

// 	messageParser := make([]string, len(kenobi))

// 	var i int
// 	for _, kenobiText := range kenobi {
// 		if kenobiText != "" {
// 			messageParser[i] = kenobiText
// 		} else if skywalker[i] != "" {
// 			messageParser[i] = skywalker[i]
// 		} else if sato[i] != "" {
// 			messageParser[i] = sato[i]
// 		}
// 		i++
// 	}
// 	fmt.Printf("\nArray Parseado: %#v", messageParser)
// }
