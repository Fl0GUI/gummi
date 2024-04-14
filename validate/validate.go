package validate

func Validate() Functions {
	f := Functions{}
	f.Sammi = ValidateSammi()
	f.Gumroad = ValidateGumroad()
	return f
}
