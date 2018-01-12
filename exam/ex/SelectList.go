package ex

import (
	"github.com/iamGreedy/gumi"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

var SelectList = []func(*gumi.Screen, gumi.Theme) testing.BenchmarkResult{
	HelloWorld,
	VLayoutCase,
	HLayoutCase,
	Boundary,
	RulerHelloWorld,
	Grid,
}

func SelectListStrings() []string {
	result := make([]string, len(SelectList))
	for i, v := range SelectList {
		temp := strings.Split(runtime.FuncForPC(reflect.ValueOf(v).Pointer()).Name(), ".")
		result[i] = temp[len(temp)-1]
	}
	return result
}
func IndexSelectString(i int) string {
	temp := strings.Split(runtime.FuncForPC(reflect.ValueOf(SelectList[i]).Pointer()).Name(), ".")
	return temp[len(temp)-1]
}
func NameSelectFunc(str string) func(*gumi.Screen, gumi.Theme) testing.BenchmarkResult {
	for i := 0; i < len(SelectList); i++ {
		if IndexSelectString(i) == str {
			return SelectList[i]
		}
	}
	return nil
}
