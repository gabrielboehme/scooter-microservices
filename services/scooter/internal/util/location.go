package util

import (
	"math"
)


func CalculateLocationMinMaxRange(latitude, longitude, radiusKm float64) (float64, float64, float64, float64) {
    // Calculate the boundaries of the range
    latitudeMin := latitude - (radiusKm / 111.32) // 1 degree of latitude is approximately 111.32 kilometers
    latitudeMax := latitude + (radiusKm / 111.32)
    longitudeMin := longitude - (radiusKm / (111.32 * math.Cos(latitude*math.Pi/180.0))) // Adjust for longitude depending on latitude
    longitudeMax := longitude + (radiusKm / (111.32 * math.Cos(latitude*math.Pi/180.0)))
    
    return latitudeMin, latitudeMax, longitudeMin, longitudeMax
}