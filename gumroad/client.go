package gumroad

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"j322.ica/gumroad-sammi/config"
)

type Client struct {
	accessToken string
	ip          string
	port        string
}

func NewClient(c *config.Config) *Client {
	conf := c.GumroadConfig
	return &Client{conf.AccessToken, conf.PublicIp, conf.ServerPort}
}

func (c *Client) Subscribe() (chan Sale, chan error) {
	p := make(chan Sale, 10)
	e := make(chan error)
	go func() {
		e <- c.setup(p)
		close(e)
	}()
	return p, e
}

func (c *Client) GetProducts() (*Products, error) {
	v := url.Values{}
	resp, err := c.doRequest(v, http.MethodGet, "https://api.gumroad.com/v2/products")
	if err != nil {
		return nil, fmt.Errorf("Could not request products: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Could not get products because of an invalid request: %s", resp.Status)
	}

	res := &Products{}
	err = json.NewDecoder(resp.Body).Decode(res)
	if err != nil {
		return nil, fmt.Errorf("Could not decode the products response: %w", err)
	}

	return res, nil
}

func (c *Client) getSubscriptions() (SubscriptionsResponse, error) {
	v := url.Values{}
	v.Set("resource_name", "sale")
	resp, err := c.doRequest(v, http.MethodGet, "https://api.gumroad.com/v2/resource_subscriptions")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println(resp.Status)
		panic(errors.New("Bad status response"))
	}

	jsonDec := json.NewDecoder(resp.Body)
	res := SubscriptionsResponse{}
	jsonDec.Decode(&res)
	return res, nil
}

func (c *Client) deleteSubscription(s Subscription) error {
	v := url.Values{}
	url := fmt.Sprintf("https://api.gumroad.com/v2/resource_subscriptions/%s", s.Id)
	resp, err := c.doRequest(v, http.MethodDelete, url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (c *Client) subscribe() error {
	v := url.Values{}
	v.Set("resource_name", "sale")
	v.Set("post_url", fmt.Sprintf("http://%s:%s%s", c.ip, c.port, listenUrl))

	resp, err := c.doRequest(v, http.MethodPut, "https://api.gumroad.com/v2/resource_subscriptions")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (c *Client) doRequest(values url.Values, method, url string) (*http.Response, error) {
	values.Set("access_token", c.accessToken)

	r, err := http.NewRequest(
		method,
		url,
		strings.NewReader(values.Encode()),
	)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(values.Encode())))

	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(r)
}
