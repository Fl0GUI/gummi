package connect

import (
	"time"

	"j322.ica/gumroad-sammi/config"
)

func backoff(action func() error, c *config.Configuration) error {
	err := action()
	for i := 0; err != nil && i < c.Advanced.BackoffAttempts; i++ {
		<-time.After(time.Duration(i) * time.Second)
		err = action()
	}
	return err
}
