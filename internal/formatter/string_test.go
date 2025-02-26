package formatter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warrenb95/cron-parser/internal/formatter"
	"github.com/warrenb95/cron-parser/internal/parser"
)

func TestFormat(t *testing.T) {
	tests := map[string]struct {
		fields   parser.Fields
		expected string
	}{
		"single values": {
			fields: parser.Fields{
				Minutes:     []int{0},
				Hours:       []int{1},
				DaysOfMonth: []int{2},
				Months:      []int{3},
				DaysOfWeek:  []int{4},
				Command:     "command",
			},
			expected: "minute        0\n" +
				"hour          1\n" +
				"day of month  2\n" +
				"month         3\n" +
				"day of week   4\n" +
				"command       command",
		},

		"multiple values": {
			fields: parser.Fields{
				Minutes:     []int{0, 15, 30, 45},
				Hours:       []int{0, 1},
				DaysOfMonth: []int{1, 2, 3},
				Months:      []int{1, 2},
				DaysOfWeek:  []int{0, 1, 2},
				Command:     "echo",
			},
			expected: "minute        0 15 30 45\n" +
				"hour          0 1\n" +
				"day of month  1 2 3\n" +
				"month         1 2\n" +
				"day of week   0 1 2\n" +
				"command       echo",
		},

		"truncated; more than 14 values": {
			fields: parser.Fields{
				Minutes: func() []int {
					nums := make([]int, 20)
					for i := 0; i < 20; i++ {
						nums[i] = i
					}
					return nums
				}(),
				Hours:       []int{0},
				DaysOfMonth: []int{1},
				Months:      []int{1},
				DaysOfWeek:  []int{0},
				Command:     "cmd",
			},
			expected: "minute        0 1 2 3 4 5 6 7 8 9 10 11 12 13\n" +
				"hour          0\n" +
				"day of month  1\n" +
				"month         1\n" +
				"day of week   0\n" +
				"command       cmd",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := formatter.Format(test.fields)
			assert.Equal(t, test.expected, result, "formatted string output")
		})
	}
}
