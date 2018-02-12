package gumi

type Information struct {
	Dt         int64
	ThrowCache bool
}

func (s Information) Require() bool {
	return s.Dt != 0 || s.ThrowCache
}
func DefaultInformation() *Information {
	return &Information{
		Dt: 0,
	}
}
