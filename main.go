package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://www.kita.net/cmmrcInfo/ehgtGnrlzInfo/rltmEhgt.do"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	html, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	exchangeRateTable := html.Find("div.boardArea")
	rows := exchangeRateTable.Find("tr")

	rows.Each(func(idx int, sel *goquery.Selection) {
		title := sel.Find("th>a").Text()
		rate := sel.Find("td").First().Text()

		fmt.Println(title, rate)
	})
}
