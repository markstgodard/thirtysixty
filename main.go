package main

import (
	"bytes"
	"net/http"
	"os"
	"time"

	"github.com/markstgodard/thirtysixty/alerter"
	"github.com/markstgodard/thirtysixty/fetcher"
	"github.com/markstgodard/thirtysixty/scraper"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("failed to load config: %s", err)
	}

	urls := viper.GetStringSlice("urls")

	for _, v := range urls {
		log.Infof("registering %s\n", v)
	}

	// fetcher fetches product info from a site
	fetcher := fetcher.Fetcher{
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

	// scraper scrapes html
	scrape := scraper.Scraper{
		Match: "Add to cart",
	}

	// alerter if a positive match
	accountSID := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	from := os.Getenv("TWILIO_FROM_PHONE_NUMBER")
	to := os.Getenv("TWILIO_TO_PHONE_NUMBER")
	alert := alerter.NewTwilioAlerter(accountSID, authToken, from, to)

	log.Info("Starting up...")

	// run on an interval
	ticker := time.NewTicker(60 * time.Second)
	done := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:

				// TODO: currently sequential, could move to channels
				for _, url := range urls {
					// fetch
					body, err := fetcher.Fetch(url)
					if err != nil {
						log.Error(err)
						break
					}

					// scrape
					var result scraper.Result
					result, err = scrape.Scrape(bytes.NewReader(body))
					if err != nil {
						log.Error(err)
						break
					}

					// alert
					err = alert.Alert(alerter.Info{
						ProductURL: url,
						Available:  result.Available,
						Price:      result.Price,
					})
					if err != nil {
						log.Error(err)
						break
					}
				}

			case <-done:
				ticker.Stop()
				return
			}
		}
	}()

	<-done
}
