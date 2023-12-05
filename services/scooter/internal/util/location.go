package util

import (
	"math"
)

const EarthRadius = 6371 // Earth's radius in kilometers


func CalculateLocationMinMaxRange(latitude, longitude, radiusKm float64) (float64, float64, float64, float64) {
    // Calculate the boundaries of the range
    latitudeMin := latitude - (radiusKm / 111.32) // 1 degree of latitude is approximately 111.32 kilometers
    latitudeMax := latitude + (radiusKm / 111.32)
    longitudeMin := longitude - (radiusKm / (111.32 * math.Cos(latitude*math.Pi/180.0))) // Adjust for longitude depending on latitude
    longitudeMax := longitude + (radiusKm / (111.32 * math.Cos(latitude*math.Pi/180.0)))
    
    return latitudeMin, latitudeMax, longitudeMin, longitudeMax
}

func DistanceTwoLocations(lat1, lon1, lat2, lon2 float64) float64 {
    // Convert latitude and longitude from degrees to radians
    lat1Rad := lat1 * (math.Pi / 180)
    lon1Rad := lon1 * (math.Pi / 180)
    lat2Rad := lat2 * (math.Pi / 180)
    lon2Rad := lon2 * (math.Pi / 180)

    // Haversine formula
    dLat := lat2Rad - lat1Rad
    dLon := lon2Rad - lon1Rad
    a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(dLon/2)*math.Sin(dLon/2)
    c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
    distance := EarthRadius * c

    return distance
}