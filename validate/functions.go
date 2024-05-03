package validate

type Functions struct {
	Sammi      error
	Gumroad    error
	FourthWall error
	Throne     error
}

func (f *Functions) Valid() bool {
	return f.Sammi == nil && f.Gumroad == nil && f.FourthWall == nil && f.Throne == nil
}

func (f *Functions) Unwrap() []error {
	return []error{f.Sammi, f.Gumroad, f.FourthWall, f.Throne}
}
