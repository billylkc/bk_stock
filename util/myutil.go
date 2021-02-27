package util

import (
	"encoding/json"
	"fmt"
)

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

// PercentChange calculates the percentage changes
func PercentChange(current, prev float64) float64 {
	if current == 0 || prev == 0 {
		return 0
	}
	percent := (current/prev - 1) * 100
	return percent
}

// PercentFormat changes the float to 1 d.p
// and add +/- sign to it
func PercentFormat(input float64) string {
	result := fmt.Sprintf("%.1f", input)
	if input >= 0 {
		result = "+" + result + "%"
	} else {
		result = result + "%" // No need additional negative sign
	}
	return result
}
