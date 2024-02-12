package unittest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	result := SumOfTwo(10, 10)
	if result != 20 {
		panic("result should be 20")
	}
}

func TestTableSum(t *testing.T) {
	tests := []struct {
		request  int
		expected int
		errMsg   string
	}{
		{
			request:  SumOfTwo(10, 10),
			expected: 20,
			errMsg:   "Result has to be 20!",
		},
		{
			request:  SumOfTwo(25, 25),
			expected: 50,
			errMsg:   "Result has to be 50!",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			require.Equal(t, test.expected, test.request, test.errMsg)
		})
	}
}

// Fail (gagal tapi flow jalan), FailNow (gagal dan flow berhenti), Error (Fail + log), Fatal (FailNow + log)
// Testify -> framework unit test go -> assert (fail), require (failnow)
