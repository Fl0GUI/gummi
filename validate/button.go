package validate

import (
	"j322.ica/gumroad-sammi/config"
	"j322.ica/gumroad-sammi/sammi"
)

func validateButton(config *config.Configuration, buttonId string) error {
	bc := sammi.NewButtonClient(&config.SammiConfig, buttonId)
	return bc.Ping()
}
