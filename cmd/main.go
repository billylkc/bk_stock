package main

import (
	"encoding/json"
	"fmt"

	"github.com/billylkc/stock"
)

func main() {
	res, err := stock.GetStockPrice(5)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(PrettyPrint(res))

}

// func main() {
// 	var (
// 		date string // date in yyyy-mm-dd format
// 	)
// 	flag.StringVar(&date, "d", "", "date in yyyy-mm-dd format")
// 	flag.Parse()

// 	if date == "" {
// 		date = time.Now().Format("2006-01-02") // default for today
// 	}

// 	// getSector(date)
// 	getIndustry(date)
// }

// func getSector(date string) {
// 	res, err := stock.GetSectorOveriew(date)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	fmt.Println(PrettyPrint(res))
// }

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

// func getIndustry(date string) {

// 	res, err := stock.GetIndustryOverview(date)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	fmt.Println(PrettyPrint(res))
// }
