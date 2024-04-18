package validate

import (
	"errors"
	"time"

	"j322.ica/gumroad-sammi/config"
	"j322.ica/gumroad-sammi/fourthwall"
)

type testMode struct {
	TestMode bool
}

func ValidateFourthWall(config *config.Configuration) error {
	if !config.FourthWallConfig.Active {
		return nil
	}

	if err := validateButton(config, config.FourthWallConfig.ButtonId); err != nil {
		return err
	}
	return nil
}

func ValidateFourthWallHook(config *config.Configuration) error {
	sales := fourthwall.GetSalesChan()

	timeout := time.After(time.Minute)
	for {
		select {
		case s := <-sales:
			value, ok := s["testMode"]
			if !ok {
				return errors.New("Invalid sale format")
			}
			testMode, ok := value.(bool)
			if !ok {
				return errors.New("Invalid sale format")
			}
			if !testMode {
				return errors.New("Expected to receive a test sale, got a real one instead.")
			}
			return nil

			break
		case <-timeout:
			return errors.New("No sale was got")
		}
	}

	return nil
}

func flushSales(salesChan chan fourthwall.Sale) {

	for {
		select {
		case <-salesChan:
			continue
		default:
			return
		}
	}
}
