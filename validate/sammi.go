package validate

import (
	"j322.ica/gumroad-sammi/config"
	"j322.ica/gumroad-sammi/sammi"
)

func ValidateSammi() error {
	testClient := sammi.NewClient(&config.Config.SammiConfig)
	return testClient.Ping()
}
