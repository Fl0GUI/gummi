package gumroad

import "net/url"

type SubscriptionsResponse struct {
	Success       bool           `json:"success"`
	Subscriptions []Subscription `json:"resource_subscriptions"`
}

type SubscriptionResponse struct {
	Success      bool         `json:"success"`
	Subscription Subscription `json:"resource_subscription"`
}

type Subscription struct {
	Id      string `json:"id"`
	Name    string `json:"resource_name"`
	PostUrl string `json:"post_url"`
}

type Products struct {
	Success  bool       `json:"success"`
	Products []Products `json:"products"`
}

type Product struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Sale url.Values
