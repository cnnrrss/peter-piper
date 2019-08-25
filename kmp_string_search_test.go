package peterpiper

import (
	"reflect"
	"testing"
)

// tcs is a slice of testcases
var tcs = []struct {
	textToSearch string
	subText      string
	expected     []int
}{
	// positive test cases
	{
		"Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!",
		"Peter",
		[]int{0, 19, 74},
	},
	{
		"Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!",
		"peter",
		[]int{0, 19, 74},
	},
	{
		"Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!",
		"pick",
		[]int{29, 57},
	},
	{
		"Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!",
		"pi",
		[]int{29, 36, 42, 50, 57},
	},
	{
		"Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!",
		"!",
		[]int{91},
	},
	{
		"",
		"",
		[]int{},
	},
	// negative test cases
	{
		"Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!",
		"z",
		[]int{},
	},
	{
		"Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!",
		"xxxxxxx",
		[]int{},
	},
	{
		"Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!",
		"#",
		[]int{},
	},
}

// TestKMP executges unit test cases for the KMP function
func TestKMPStringSearch(t *testing.T) {
	for _, tc := range tcs {
		t.Logf("~ %v ~\n", tc)
		actual := KMPStringSearch(tc.textToSearch, tc.subText)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("\nTextToSearch: \"%s\"\nSubtext:\"%s\"\nResult incorrect got: %v, want: %v.", tc.textToSearch, tc.subText, actual, tc.expected)
			t.Fail()
		}
	}
}

func Benchmark_KMPStringSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			KMPStringSearch(tc.textToSearch, tc.subText)
		}
	}
}
