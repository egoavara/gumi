package sdl2example

import (
	"github.com/iamGreedy/gumi"
	"runtime"
	"reflect"
)

var kinds = []gumi.GUMI{
	HelloWorld,
	HelloButton,
}

func List() []string {
	var temp []string
	for _, v := range kinds{
		temp = append(temp, runtime.FuncForPC(reflect.ValueOf(v).Pointer()).Name())
	}
	return temp
}
func Select(name string) gumi.GUMI{
	for _, v := range kinds{
		if name == runtime.FuncForPC(reflect.ValueOf(v).Pointer()).Name(){
			return v
		}
	}
	return nil
}
