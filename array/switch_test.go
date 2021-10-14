package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwitchType(t *testing.T) {
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
			name:       "arrayIntToArrayString",
			args:       args{what: []int{3, 2}, where: []string{"13", "2", "3"}},
			wantResult: true,
		},
		{
			name:       "arrayIntToArrayInt",
			args:       args{what: []int{3, 2}, where: []int{1, 2}},
			wantResult: false,
		},
		{
			name:       "arrayString",
			args:       args{what: []string{"3", "13"}, where: []string{"13", "14", "3"}},
			wantResult: false, // TODO: change after adding functional
		},
		{
			name:       "arrayStringInUndefinedType",
			args:       args{what: func() {}, where: []string{"3", "4"}},
			wantResult: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SwitchType(tt.args.what, tt.args.where)
			assert.Equal(t, tt.wantResult, result)
		})
	}
}
