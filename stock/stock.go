package stock

import (
	"fmt"
	"time"

	"github.com/billylkc/stock/db"
	"github.com/billylkc/stock/util"
)

type StockPrice struct {
	Code     string
	DateRaw  time.Time // real date format
	Date     string    // date in string format, DD/MM
	Close    float64
	Changes  float64 // Percentage changes in float
	ChangesF string  // Percentage changes on Close. Formatted with +/- sign
}

// RecordExists checks if we have records for a particular date in the database
// true if already exists
func RecordExists(date string) bool {
	db, err := db.GetConnection()
	if err != nil {
		return true
	}
	queryF := `
    SELECT count(1) as cnt
    FROM stock
    WHERE date = '%s'`

	query := fmt.Sprintf(queryF, date)
	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	var num int
	for rows.Next() {
		_ = rows.Scan(&num)
	}
	if num > 0 {
		return true
	}
	return false // false as safe to insert
}

// GetStockPrice gets the historical stock price of a certain code
func GetStockPrice(code int) ([]StockPrice, error) {
	var result []StockPrice

	db, err := db.GetConnection()
	if err != nil {
		return result, err
	}

	// Query data
	c := fmt.Sprintf("%05d", code)
	queryF := `
    SELECT
       code, date, close
    FROM
       stock
    WHERE
       code = '%s'
    ORDER BY
       date desc
    `
	query := fmt.Sprintf(queryF, c)
	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	for rows.Next() {
		var sp StockPrice
		_ = rows.Scan(&sp.Code, &sp.DateRaw, &sp.Close)
		sp.Date = sp.DateRaw.Format("02/01")
		result = append(result, sp)
	}

	// Derive % changes on Close
	for i, _ := range result {
		var changes float64
		if i < len(result)-1 {
			changes = util.PercentChange(result[i].Close, result[i+1].Close)
		}
		result[i].Changes = changes
		result[i].ChangesF = util.PercentFormat(changes)
	}

	return result, nil
}
