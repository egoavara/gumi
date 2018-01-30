package gumi


type Information struct {
	Dt uint64
}

func DefaultInformation() *Information {
	return &Information{
		Dt: 0,
	}
}