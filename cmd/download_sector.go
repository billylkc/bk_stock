package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/billylkc/stock/stock"
)

func main() {
	var (
		date string // date in yyyy-mm-dd format
	)
	flag.StringVar(&date, "d", "", "date in yyyy-mm-dd format")
	flag.Parse()

	if date == "" {
		date = time.Now().Format("2006-01-02") // default for today
	}

	getSector(date)
}

func getSector(date string) {
	res, err := stock.GetSectorOveriew(date)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = stock.InsertSector(res)
	if err != nil {
		fmt.Println(err.Error())
	}

}
