package ex

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

//
func PrintResult(result testing.BenchmarkResult) {
	fmt.Printf("%-50s : %d\n", "Total Run", result.N)
	fmt.Printf("%-50s : %s\n", "Total Duration", result.T.String())
	fmt.Printf("%-50s : %s\n", "Total Memory Allocation", FormatBytes(result.MemAllocs))
	fmt.Printf("%-50s : %s\n", "Total Memory Bytes", FormatBytes(result.MemBytes))
	fmt.Println()
	fmt.Printf("%-50s : %dns\n", "Duration per OP", result.NsPerOp())
	fmt.Printf("%-50s : %d\n", "Allocated Bytes per OP", result.AllocedBytesPerOp())
	fmt.Printf("%-50s : %d\n", "Allocation per OP", result.AllocsPerOp())
}
func FormatBytes(int uint64) string {
	// Tib
	if int > Pow(1024, 4) {
		temp := math.Floor(float64(int) / float64(Pow(1024, 3)))
		return strconv.FormatFloat(temp/1000., 'f', 4, 64) + "Tib"
	}
	// Gib
	if int > Pow(1024, 3) {
		temp := math.Floor(float64(int) / float64(Pow(1024, 2)))
		return strconv.FormatFloat(temp/1000., 'f', 4, 64) + "Gib"
	}
	// Mib
	if int > Pow(1024, 2) {
		temp := math.Floor(float64(int) / float64(Pow(1024, 1)))
		return strconv.FormatFloat(temp/1000., 'f', 4, 64) + "Mib"
	}
	// Kib
	if int > Pow(1024, 1) {
		temp := math.Floor(float64(int))
		return strconv.FormatFloat(temp/1000., 'f', 4, 64) + "Kib"
	}
	return strconv.FormatUint(int, 10) + "b"
}

// x ^ n
func Pow(x uint64, n uint64) uint64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}
	result := uint64(1)
	for i := uint64(0); i < n; i++ {
		result = result * x
	}
	return result

}

//
func StringsToString(strs ...string) string {
	result := "<"
	for _, v := range strs {
		result += v
		result += "/"
	}
	result = result[:len(result)-1]
	result += ">"
	return result
}
