package validate

import (
	"j322.ica/gumroad-sammi/config"
	"j322.ica/gumroad-sammi/gumroad"
)

func ValidateGumroad(c *config.Configuration) error {
	if !c.GumroadConfig.Active {
		return nil
	}

	if err := validateServer(c); err != nil {
		return err
	}
	if err := validateAccess(c); err != nil {
		return err
	}
	return nil
}

func validateServer(config *config.Configuration) error {
	c := gumroad.NewClient(config)
	return c.Ping()
}

func validateAccess(config *config.Configuration) error {
	c := gumroad.NewClient(config)
	_, err := c.GetProducts()
	return err
}
