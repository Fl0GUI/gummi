package connect

import (
	"time"

	"j322.ica/gumroad-sammi/config"
)

func backoff(action func() error, c *config.BackoffConfig) error {
	err := action()
	for i := 0; err != nil && (i < c.Max || c.Max == 0); i++ {
		<-time.After(time.Duration(c.Increment*i)*time.Second + time.Duration(c.Base)*time.Second)
		err = action()
	}
	return err
}
