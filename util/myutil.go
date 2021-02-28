package util

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
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

// ParseF parses from string to float, ignoring % sign
func ParseF(s string) (float64, error) {
	s = strings.ReplaceAll(s, "%", "")
	num, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0, nil
	}
	return num, err
}

// ParseI parses from string to integer, also handles numbers like 1K, 1M, 1B, etc
func ParseI(s string) (int, error) {
	var (
		num int
		f   float64
		err error
	)

	if strings.Contains(s, "N/A") {
		num = 0
	}
	if strings.Contains(s, "K") {
		s = strings.ReplaceAll(s, "K", "")
		f, err = strconv.ParseFloat(s, 64)
		num = int(f * 1_000)
	}
	if strings.Contains(s, "M") {
		s = strings.ReplaceAll(s, "M", "")
		f, err = strconv.ParseFloat(s, 64)
		num = int(f * 1_000_000)
	}
	if strings.Contains(s, "B") {
		s = strings.ReplaceAll(s, "B", "")
		f, err = strconv.ParseFloat(s, 64)
		num = int(f * 1_000_000_000)
	}
	if err != nil {
		return 0, err
	}

	return num, nil
}
