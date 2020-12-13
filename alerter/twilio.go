package alerter

import (
	"fmt"

	twilio "github.com/kevinburke/twilio-go"

	log "github.com/sirupsen/logrus"
)

type TwilioAlerter struct {
	from   string
	to     string
	client *twilio.Client
}

func NewTwilioAlerter(accountSID, authToken, from, to string) *TwilioAlerter {
	client := twilio.NewClient(accountSID, authToken, nil)
	return &TwilioAlerter{
		from:   from,
		to:     to,
		client: client,
	}
}

func (t *TwilioAlerter) Alert(info Info) error {

	log.WithFields(log.Fields{
		"product":   info.ProductURL,
		"available": info.Available,
		"price":     info.Price,
	}).Info("alert for product")

	// if available send an SMS
	if info.Available {
		say := fmt.Sprintf("Product now available: %f %s", info.Price, info.ProductURL)

		msg, err := t.client.Messages.SendMessage(t.from, t.to, say, nil)
		if err != nil {
			log.Error(err)
		}

		fmt.Printf("msg: %#v\n", msg)

	}

	return nil
}
