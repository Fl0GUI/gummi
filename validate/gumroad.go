package validate

import (
	"j322.ica/gumroad-sammi/config"
	"j322.ica/gumroad-sammi/gumroad"
	"j322.ica/gumroad-sammi/sammi"
)

func ValidateGumroad() error {
	if err := validateServer(); err != nil {
		return err
	}
	if err := validateButton(); err != nil {
		return err
	}
	if err := validateAccess(); err != nil {
		return err
	}
	return nil
}

func validateServer() error {
	c := gumroad.NewClient()
	return c.Ping()
}

func validateAccess() error {
	c := gumroad.NewClient()
	_, err := c.GetProducts()
	return err
}

func validateButton() error {
	bc := sammi.NewButtonClient(&config.Config.SammiConfig, config.Config.GumroadConfig.ButtonId)
	return bc.Ping()
}
