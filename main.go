package main

import (
	"fmt"

	"github.com/coseo12/gps-calculate/converter"
	"github.com/coseo12/gps-calculate/parser"
)

func main() {
	convertData := converter.ConvertData{Lat: 37.5665, Lng: 126.9780, X: 1.59, Y: 2.91}
	lat, lng := converter.GetConvert(convertData)

	fmt.Printf("%v, %v \n", lat, lng)

	data := parser.GetJson()

	fmt.Println(data)
}
