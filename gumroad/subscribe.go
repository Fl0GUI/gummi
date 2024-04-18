package gumroad

import (
	"errors"
	"fmt"
)

func (c *Client) Unsubscribe() error {
	// check subscriptions
	subscriptions, err := c.getSubscriptions()
	if err != nil {
		return fmt.Errorf("I could not get my old gumroad subscriptions: %w", err)
	}

	// delete old subscriptions
	errs := make([]error, 0)
	for _, subscription := range subscriptions.Subscriptions {
		if subscription.PostUrl == c.subscriptionURL() {
			err = c.deleteSubscription(subscription)
			if err != nil {
				errs = append(errs, fmt.Errorf("Could not delete an old gumroad subscription: %w", err))
			}
		}
	}
	return errors.Join(errs...)
}

func (c *Client) Subscribe() error {
	if err := c.Unsubscribe(); err != nil {
		return fmt.Errorf("Could not delete old gumroad subscriptions: %w", err)
	}

	if err := c.subscribe(); err != nil {
		return fmt.Errorf("I could not make a new gumroad subscription: %w", err)
	}

	return nil
}
