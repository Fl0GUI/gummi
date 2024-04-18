package validate

import (
	"j322.ica/gumroad-sammi/config"
	"j322.ica/gumroad-sammi/sammi"
)

func ValidateSammi(config *config.Configuration) error {
	testClient := sammi.NewClient(&config.SammiConfig)
	return testClient.Ping()
}
