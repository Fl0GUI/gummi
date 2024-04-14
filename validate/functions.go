package validate

type Functions struct {
	Sammi   error
	Gumroad error
}

func (f *Functions) Valid() bool {
	return f.Sammi == nil && f.Gumroad == nil
}

func (f *Functions) Unwrap() []error {
	return []error{f.Sammi, f.Gumroad}
}
