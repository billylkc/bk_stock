package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	dev()
}

func dev() {
	fmt.Println("Dev")
}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
