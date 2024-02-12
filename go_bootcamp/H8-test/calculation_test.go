package h8test

import (
	"testing"
	// "github.com/RichardKnop/machinery/v2/backends/result"
)

func TestSum(t *testing.T) {
	result := Sum(10, 10)

	if result != 20 {
		panic("Rseult should be 20")
	}
}
