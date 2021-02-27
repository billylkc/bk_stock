package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/billylkc/stock/quandl"
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

	exists := stock.RecordExists(date)
	if exists {
		fmt.Printf("records exists - %s\n", date)
		os.Exit(0)
	}

	q := quandl.New()
	res, err := q.GetStockByDate(date)
	if err != nil {
		fmt.Println(err)
	}
	err = q.Insert(res)
	if err != nil {
		fmt.Println(err.Error())
	}

}
