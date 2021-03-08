package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/billylkc/stock/quandl"
	"github.com/billylkc/stock/stock"
	"github.com/sirupsen/logrus"
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

	// Setup logger
	file, err := os.OpenFile("/var/log/stock.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	mw := io.MultiWriter(file, os.Stdout)
	logger := logrus.New()
	logger.Out = mw
	logger.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05", FullTimestamp: true})

	q := quandl.New(logger)
	logger.Info(fmt.Sprintf("Start getting stock - %s", date))
	res, err := q.GetStockByDate(date)
	if err != nil {
		logger.Error(err)
	}
	err = q.Insert(res)
	if err != nil {
		logger.Error(err)
	}

}
