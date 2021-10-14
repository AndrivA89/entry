package single

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
			name:       "singleIntToArrayString",
			args:       args{what: 3, where: []string{"13", "14", "3"}},
			wantResult: true,
		},
		{
			name:       "singleIntToArrayInt",
			args:       args{what: 3, where: []int{13, 14, 3}},
			wantResult: true,
		},
		{
			name:       "singleStringToArrayString",
			args:       args{what: "3", where: []string{"13", "14", "3"}},
			wantResult: true,
		},
		{
			name:       "singleStringToArrayInt",
			args:       args{what: "3", where: []int{13, 14, 3}},
			wantResult: true,
		},
		{
			name:       "singleStringInUndefinedType",
			args:       args{what: func() {}, where: "3"},
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
