package alerter

import (
	log "github.com/sirupsen/logrus"
)

type ConsoleAlerter struct {
}

func NewConsoleAlerter() *ConsoleAlerter {
	return &ConsoleAlerter{}
}

func (c *ConsoleAlerter) Alert(info Info) error {

	log.WithFields(log.Fields{
		"product":   info.ProductURL,
		"available": info.Available,
		"price":     info.Price,
	}).Info("alert for product")
	return nil
}
