package task0001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name string
		args struct {
			nums   []int
			target int
		}
		want []int
	}{
		{
			name: "testOne",
			args: struct {
				nums   []int
				target int
			}{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, twoSum(tt.args.nums, tt.args.target))
		})
	}

}
