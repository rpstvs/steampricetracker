package utils

import (
	"strings"
	"time"
)

func ConvertDate() time.Time {
	currentTime := time.Now()

	// Truncate to get only the date part
	currentDate := currentTime.Truncate(24 * time.Hour)

	return currentDate
}

func ExtractSteamid(steamid string) string {
	tmp := strings.Split(steamid, "/")
	id := tmp[len(tmp)-1]
	return id
}

func BuildImageURL(imageId string) string {
	tmp := "https://community.akamai.steamstatic.com/economy/image/"

	return tmp + imageId
}

func ParsePrice(price string) float64 {

	return 0.00
}
