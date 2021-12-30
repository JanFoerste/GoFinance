package library

import (
	"fmt"
	"math/rand"
	"time"
)

type RGBColor struct {
	Red   int
	Green int
	Blue  int
}

func convertToHex(number int) string {
	hex := fmt.Sprintf("%x", number)
	if len(hex) == 1 {
		hex = "0" + hex
	}
	return hex
}

func GetRandomRGBColor() RGBColor {
	rand.Seed(time.Now().UnixNano())

	Red := rand.Intn(255)
	Green := rand.Intn(255)
	Blue := rand.Intn(255)
	color := RGBColor{Red, Green, Blue}

	return color
}

func GetRandomHexColor() string {
	color := GetRandomRGBColor()
	hex := "#" + convertToHex(color.Red) + convertToHex(color.Green) + convertToHex(color.Blue)

	return hex
}
