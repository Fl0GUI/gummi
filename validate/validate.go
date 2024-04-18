package validate

import (
	"j322.ica/gumroad-sammi/config"
)

func UpdateValidation(f *Functions, c *config.Configuration) {
	f.Sammi = ValidateSammi(c)
	f.Gumroad = ValidateGumroad(c)
	f.FourthWall = ValidateFourthWall(c)
}

func Validate(c *config.Configuration) Functions {
	f := Functions{}
	UpdateValidation(&f, c)
	return f
}
