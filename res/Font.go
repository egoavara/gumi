package res

import (
	"runtime"
	"path/filepath"
	"io/ioutil"
)

var (
	NanumSquareRoundB  []byte
	NanumSquareRoundEB []byte
	NanumSquareRoundL  []byte
	NanumSquareRoundR  []byte
)

func init() {
	_, p, _, ok := runtime.Caller(0)
	if !ok{
		panic('W')
	}
	d := filepath.Dir(p)
	var err error

	NanumSquareRoundL, err = ioutil.ReadFile(filepath.Join(d, "NanumSquareRoundL.ttf"))
	if err != nil {
		panic(err)
	}
	NanumSquareRoundR, err = ioutil.ReadFile(filepath.Join(d, "NanumSquareRoundR.ttf"))
	if err != nil {
		panic(err)
	}
	NanumSquareRoundB, err = ioutil.ReadFile(filepath.Join(d, "NanumSquareRoundB.ttf"))
	if err != nil {
		panic(err)
	}
	NanumSquareRoundEB, err = ioutil.ReadFile(filepath.Join(d, "NanumSquareRoundEB.ttf"))
	if err != nil {
		panic(err)
	}

}
