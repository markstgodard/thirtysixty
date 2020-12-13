package scraper

import (
	"io"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
)

type Result struct {
	Price     float64
	Available bool
}

type Scraper struct {
	Match string
}

func (s *Scraper) Scrape(body io.Reader) (Result, error) {
	var res Result

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return res, err
	}

	// price
	doc.Find("#app > div.page-content > div.page-section > div > div > div.row-side > div.product-buy-box > div.product-pane > div.product-price > ul > li.price-current").Each(func(i int, sel *goquery.Selection) {
		v := strings.TrimSpace(sel.Find("strong").Text())
		v = strings.Replace(v, ",", "", -1)

		price, err := strconv.ParseFloat(v, 64)
		if err != nil {
			log.Errorf("failed to parse price: %s\n", v)
		}
		res.Price = price
	})

	// if available to add to cart
	doc.Find("#ProductBuy > div > div:nth-child(2)").Each(func(i int, sel *goquery.Selection) {
		status := strings.TrimSpace(sel.Find("button").Text())
		if strings.EqualFold(s.Match, status) {
			res.Available = true
		}
	})

	return res, nil
}
