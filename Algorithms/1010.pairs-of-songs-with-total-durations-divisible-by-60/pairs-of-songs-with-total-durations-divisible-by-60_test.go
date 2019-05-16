package problem1010

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	time []int
	ans  int
}{

	{
		[]int{336, 24, 100, 342, 274, 11, 43, 22, 416, 138, 384, 386, 70, 265, 59, 253, 344, 435, 400, 296, 192, 143, 311, 424, 315, 63, 420, 254, 493, 431, 32, 394, 178, 51, 378, 335, 265, 92, 335, 325, 25, 355, 258, 298, 390, 399, 393, 114, 149, 62, 299, 471, 286, 204, 163, 214, 15, 272, 315, 212, 272, 437, 339, 193, 125, 394, 62, 188, 154, 150, 109, 294, 228, 200, 459, 42, 469, 132, 37, 460, 143, 1, 144, 127, 398, 82, 370, 464, 14, 85, 321, 358, 205, 14, 264, 289, 183, 93, 56, 126, 413, 140, 441, 446, 445, 378, 258, 119, 385, 226, 8, 93, 476, 265, 115, 86, 360, 92, 396, 407, 458, 58, 65, 397, 381, 32, 228, 37, 319, 220, 73, 328, 162, 458, 231, 219, 481, 387, 423, 256, 252, 36, 309, 395, 471, 4, 225, 146, 188, 182, 347, 82, 21, 292, 91, 144, 387, 263, 206, 452, 197, 192, 324, 257, 370, 28, 440, 180, 294},
		245,
	},

	{
		[]int{30, 20, 150, 100, 40},
		3,
	},

	{
		[]int{60, 60, 60},
		3,
	},

	// 可以有多个 testcase
}

func Test_numPairsDivisibleBy60(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		sort.Ints(tc.time)
		ast.Equal(tc.ans, numPairsDivisibleBy60(tc.time), "输入:%v", tc)
	}
}

func Benchmark_numPairsDivisibleBy60(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numPairsDivisibleBy60(tc.time)
		}
	}
}