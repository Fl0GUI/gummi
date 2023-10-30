package gummi

import (
	"errors"
	"log"
	"net/http"

	"j322.ica/gumroad-sammi/config"
	"j322.ica/gumroad-sammi/gumroad"
	"j322.ica/gumroad-sammi/sammi"
)

type functions struct {
	sammi         bool
	sammiAuth     bool
	sammiButton   bool
	gumroadServer bool
	gumroadToken  bool
}

func (f *functions) valid() bool {
	return f.validSammi() && f.validGumroad()
}

func (f *functions) validSammi() bool {
	return f.sammi && f.sammiAuth && f.sammiButton
}

func (f *functions) validGumroad() bool {
	return f.gumroadServer && f.gumroadToken
}

func testConfig(c *config.Config, f *functions) {
	if c == nil {
		c = &config.Config{}
	}

	testSammi(c, f)
	testGumroad(c, f)
}

func testSammi(c *config.Config, f *functions) {
	if f.validSammi() {
		return
	}
	sammiC := sammi.NewClient(c)
	err := sammiC.Ping()

	if err == nil {
		f.sammi = true
		f.sammiAuth = true
		f.sammiButton = true
	} else if errors.Is(err, sammi.AuthError) {
		f.sammi = true
		f.sammiAuth = false
	} else if errors.Is(err, sammi.ButtonIdNotFoundError) {
		f.sammi = true
		f.sammiAuth = true
		f.sammiButton = false
	} else {
		f.sammi = false
		f.sammiAuth = false
		f.sammiButton = false
	}
}

func testGumroad(c *config.Config, f *functions) {
	if f.validGumroad() {
		return
	}
	gumroadC := gumroad.NewClient(c)

	// server reachability
	dummyChan := make(chan gumroad.Sale)
	gumroadS := gumroadC.Listen(dummyChan)

	go func() {
		if err := gumroadS.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	errChan := make(chan error)
	go func() {
		errChan <- gumroadC.Ping()
	}()
	err := <-errChan

	if err := gumroadS.Close(); err != nil {
		log.Fatal(err)
	}
	f.gumroadServer = err == nil

	// access token for the api
	_, err = gumroadC.GetProducts()
	f.gumroadToken = err == nil
}
