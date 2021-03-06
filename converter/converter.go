package converter

import (
	"math"
)

const (
	R       = 6371.009		// Earth radius (unit: km)
	PI      = math.Pi		// PI
	R_ROUND = 2 * PI * R	// Earth round (unit: km)
)

type ConvertData struct {
	Lat float64
	Lng float64
	X   float64
	Y   float64
}

// Get time to distance for latitude 
func getDistanceLatitude() (float64, float64, float64) {
	lat_km_degrees := R_ROUND / 360
	lat_km_minutes := lat_km_degrees / 60
	lat_km_seconds := lat_km_minutes / 60

	return lat_km_degrees, lat_km_minutes, lat_km_seconds
}

//Get time to distance for longitude 
func getDistanceLongitude(degrees float64) (float64, float64, float64) {
	r_round_lat := math.Cos(degrees) * R_ROUND 
	lng_km_degrees := r_round_lat / 360
	lng_km_minutes := lng_km_degrees / 60
	lng_km_seconds := lng_km_minutes / 60

	return lng_km_degrees, lng_km_minutes, lng_km_seconds
}

// Second per distance
func convertMeterToSec(km float64, km_sec float64) float64 {
	return km / km_sec
}

// Converter decimal degree to degree minute second
func convertDDToDMS(dd float64) (float64, float64, float64) {
	degrees := math.Floor(dd)
	minutes := (dd - degrees) * 60
	seconds := (minutes - math.Floor(minutes)) * 60
	return degrees, minutes, seconds
}

// Converter degree minute second to decimal degree
func convertDMSToDD(degrees float64, minutes float64, seconds float64) float64 {
	return degrees + (math.Floor(minutes) / 60) + (float64(int(seconds*100)) / 100 / 3600)
}

// Get convert to data
func GetConvert(convertData ConvertData) (float64, float64) {
	// Latitude to km
	_, _, lat_km_seconds := getDistanceLatitude()
	_, _, lng_km_seconds := getDistanceLongitude(math.Floor(convertData.Lat))

	// Meter to sec
	y_sec := convertMeterToSec(convertData.Y, lat_km_seconds)
	x_sec := convertMeterToSec(convertData.X, lng_km_seconds)

	// DD to DMS
	lat_degrees, lat_minutes, lat_seconds := convertDDToDMS(convertData.Lat)
	lng_degrees, lng_minutes, lng_seconds := convertDDToDMS(convertData.Lng)

	// Sum to sec
	lat_seconds_sum := lat_seconds + y_sec
	lng_seconds_sum := lng_seconds + x_sec

	return convertDMSToDD(lat_degrees, lat_minutes, lat_seconds_sum), convertDMSToDD(lng_degrees, lng_minutes, lng_seconds_sum)
}

func ConvertDistanceInPixelsToMeter(dp float64) float64 {
	return dp / 100 / 1000 * -1
}