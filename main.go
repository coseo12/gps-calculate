package main

import (
	"fmt"

	"github.com/coseo12/gps-calculate/converter"
)

const (
	lat = 35.878925
	lng = 128.575591
) 


func main() {
	scale:= 24.248537706402594
	x := 13011.626953125 * scale / 1000 / 1000
	y := 5233.927734375 * scale / 1000 / 1000
	convertData := converter.ConvertData{Lat: lat, Lng: lng, X: x, Y: y}
	lat, lng := converter.GetConvert(convertData)

	fmt.Printf("%v, %v \n", x, y)
	fmt.Printf("%v, %v \n", lat, lng)
}