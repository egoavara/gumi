package gumre

func Assert(err error) {
	if err != nil {
		panic(err)
	}
}