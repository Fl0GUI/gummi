package validate

import (
	"j322.ica/gumroad-sammi/config"
	"j322.ica/gumroad-sammi/throne"
)

func ValidateThrone(c *config.Configuration) error {
	if !c.ThroneConfig.Active {
		return nil
	}

	if err := throne.Start(&c.ThroneConfig); err != nil {
		return err
	} else {
		throne.Stop()
		return nil
	}
}
