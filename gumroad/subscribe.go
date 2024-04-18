package gumroad

import (
	"fmt"
	"log"
)

func (c *Client) Subscribe() error {
	// check subscriptions
	subscriptions, err := c.getSubscriptions()
	if err != nil {
		log.Printf("I could not check my old subscriptions: %s\n", err)
	}

	// delete old subscriptions
	for _, subscription := range subscriptions.Subscriptions {
		err = c.deleteSubscription(subscription)
		if err != nil {
			fmt.Printf("I could not delete subcription with id %s: %s\n", subscription.Id, err)
		}
	}

	// make subscription
	err = c.subscribe()
	if err != nil {
		return fmt.Errorf("I could not make a new subscription: %w", err)
	}

	return nil
}
