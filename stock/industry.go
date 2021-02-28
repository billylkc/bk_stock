package stock

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/billylkc/stock/util"
)

// Industry details from aastock
type Industry struct {
	Sector    string
	Industry  string
	CodeF     string
	Close     float64
	Change    float64
	ChangePct float64
	Volume    int
	Turnover  int
	PE        float64 // Price per Earnings
	PB        float64 // Price to Book
	YieldPct  float64
	MarketCap int
}

// Gets all the sectors + industry code
func GetIndustryDetails(date string) ([]Industry, error) {
	var results []Industry
	links, err := getIndustryLinks(date) // check dates
	if err != nil {
		return results, err
	}

	for _, link := range links {
		industry, _ := getIndustryDetails(link)
		results = append(results, industry...)
	}
	return results, nil
}

// getIndustryDetail gets a single industry detail from aastock
func getIndustryDetails(link string) ([]Industry, error) {

	var result []Industry
	fmt.Printf("Getting link - %s\n", link)

	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Title
	var (
		sector    string // Industry sectors, e.g. Materials
		industry  string // Industry, e.g. Chemical Products
		code      string
		close     float64
		change    float64
		changePct float64
		volume    int
		turnover  int
		pe        float64
		pb        float64
		yield     float64
		marketCap int
	)
	doc.Find("h1").Each(func(i int, s *goquery.Selection) {
		text := s.Text() // e.g. Industry Details - Materials - Chemical Products
		texts := strings.Split(text, "-")
		if len(texts) == 3 {
			sector = strings.TrimSpace(texts[1])   // e.g. Materials
			industry = strings.TrimSpace(texts[2]) // e.g. Chemical Products
			fmt.Printf("Getting [%s] - [%s]\n", sector, industry)
		}
	})

	// For each code inside a sector, gets the details
	doc.Find("span.float_l").Each(func(i int, s *goquery.Selection) {
		code = strings.TrimSpace(s.Text()) // e.g. 00301.HK
		if strings.Contains(code, "0") {   // Check starts with 0
			code = strings.ReplaceAll(code, ".HK", "") // 00301.HK -> 00301
			ss := s.ParentsUntil("tbody")
			var values []string

			ss.Each(func(j int, tb *goquery.Selection) {
				tb.Find("td.cls.txt_r.pad3").Each(func(i int, td *goquery.Selection) {
					// fmt.Println(td.Text())
					values = append(values, td.Text())
				})
			})

			if len(values) == 10 {
				_ = values[0] // Some empty string
				close, _ = util.ParseF(values[1])
				change, _ = util.ParseF(values[2])
				changePct, _ = util.ParseF(values[3])
				volume, _ = util.ParseI(values[4])
				turnover, _ = util.ParseI(values[5])
				pe, _ = util.ParseF(values[6])
				pb, _ = util.ParseF(values[7])
				yield, _ = util.ParseF(values[8])
				marketCap, _ = util.ParseI(values[9])

				rec := Industry{
					Sector:    sector,
					Industry:  industry,
					CodeF:     code,
					Close:     close,
					Change:    change,
					ChangePct: changePct,
					Volume:    volume,
					Turnover:  turnover,
					PE:        pe,
					PB:        pb,
					YieldPct:  yield,
					MarketCap: marketCap,
				}
				result = append(result, rec)
			}
		}
	})

	return result, nil
}

// getIndustryLinks gets all the individual sector/industires links
func getIndustryLinks(date string) ([]string, error) {
	var links []string

	// Check if data is ready
	dataReady := checkIndustryAvailability(date)
	if !dataReady {
		return links, fmt.Errorf("data not ready - %s", date)
	}

	res, err := http.Get("http://www.aastocks.com/en/stocks/market/industry/sector-industry-details.aspx")
	if err != nil {
		return links, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return links, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}
	body, err := ioutil.ReadAll(res.Body)

	r := regexp.MustCompile("gotoindustry\\(\\'(\\d{4})\\'\\)")
	matches := r.FindAllStringSubmatch(string(body), -1)
	for _, match := range matches {
		if len(match) >= 2 {
			industryCode := match[1]
			link := fmt.Sprintf("http://www.aastocks.com/en/stocks/market/industry/sector-industry-details.aspx?industrysymbol=%s&t=1&s=&o=&p=", industryCode)

			links = append(links, link)
		}
	}
	return links, nil
}

// checkIndustryAvailability checks the date from the sector page and see if it matches the input date
func checkIndustryAvailability(date string) bool {
	res, err := http.Get("http://www.aastocks.com/en/stocks/market/industry/sector-industry-details.aspx?industrysymbol=2033&t=1&hk=0")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	body, err := ioutil.ReadAll(res.Body)
	re := regexp.MustCompile(`.*Last Update:\s*(\d{4}\/\d{2}\/\d{2})`)
	matched := re.FindAllSubmatch(body, -1)

	// TODO: better checking later
	var b bool
	if string(matched[0][1]) == date {
		b = true
	} else {
		b = false
	}
	return b
}
