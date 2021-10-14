package entry

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsContains(t *testing.T) {
	type args struct {
		what  interface{}
		where interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantResult bool
	}{
		{
			name:       "singleToArray",
			args:       args{what: 3, where: []int{13, 14, 3}},
			wantResult: true,
		},
		{
			name:       "arrayToArray",
			args:       args{what: []int{3, 2}, where: []string{"13", "14", "3", "2"}},
			wantResult: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsContains(tt.args.what, tt.args.where)
			assert.Equal(t, tt.wantResult, result)
		})
	}
}
