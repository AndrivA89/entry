package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTryConvertForInt(t *testing.T) {
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
			name:       "arrayIntInArrayString",
			args:       args{what: []int{3, 4, 5}, where: []string{"13", "14", "3", "4", "5"}},
			wantResult: true,
		},
		{
			name:       "arrayIntInArrayStringFalse",
			args:       args{what: []int{3, 4, 6}, where: []string{"13", "14", "3", "4", "5"}},
			wantResult: false,
		},
		{
			name:       "arrayIntInArrayFloat64",
			args:       args{what: []int{3, 4, 5}, where: []float64{13, 14, 3, 4, 5}},
			wantResult: true,
		},
		{
			name:       "arrayIntInArrayFloat64False",
			args:       args{what: []int{3, 4, 55}, where: []float64{13, 14, 3, 4, 5}},
			wantResult: false,
		},
		{
			name:       "arrayIntInArrayFloat32",
			args:       args{what: []int{3, 4, 5}, where: []float32{13, 14, 3, 4, 5}},
			wantResult: true,
		},
		{
			name:       "arrayIntInArrayFloat32False",
			args:       args{what: []int{3, 4, 55}, where: []float32{13, 14, 3, 4, 5}},
			wantResult: false,
		},
		{
			name:       "arrayIntInArrayUint",
			args:       args{what: []int{3, 4, 5}, where: []uint{13, 14, 3, 4, 5}},
			wantResult: true,
		},
		{
			name:       "arrayIntInArrayUintFalse",
			args:       args{what: []int{3, 4, 544}, where: []uint{13, 14, 3, 4, 5}},
			wantResult: false,
		},
		{
			name:       "arrayIntInArrayUint8",
			args:       args{what: []int{3, 4, 5}, where: []uint8{13, 14, 3, 4, 5}},
			wantResult: true,
		},
		{
			name:       "arrayIntInArrayUint8False",
			args:       args{what: []int{3, 4, 521}, where: []uint8{13, 14, 3, 4, 5}},
			wantResult: false,
		},
		{
			name:       "arrayIntInArrayUint16",
			args:       args{what: []int{3, 4, 5}, where: []uint16{13, 14, 3, 4, 5}},
			wantResult: true,
		},
		{
			name:       "arrayIntInArrayUint16False",
			args:       args{what: []int{3, 4, 511}, where: []uint16{13, 14, 3, 4, 5}},
			wantResult: false,
		},
		{
			name:       "arrayIntInArrayUint32",
			args:       args{what: []int{3, 4, 5}, where: []uint32{13, 14, 3, 4, 5}},
			wantResult: true,
		},
		{
			name:       "arrayIntInArrayUint32False",
			args:       args{what: []int{3, 4, 6}, where: []uint32{13, 14, 3, 4, 5}},
			wantResult: false,
		},
		{
			name:       "arrayIntInArrayUint64",
			args:       args{what: []int{3, 4, 5}, where: []uint64{13, 14, 3, 4, 5}},
			wantResult: true,
		},
		{
			name:       "arrayIntInArrayUint64False",
			args:       args{what: []int{3, 4, 6}, where: []uint64{13, 14, 3, 4, 5}},
			wantResult: false,
		},
		{
			name:       "arrayIntInUndefinedType",
			args:       args{what: []int{3, 4}, where: func() {}},
			wantResult: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tryConvertForInt(tt.args.what, tt.args.where)
			assert.Equal(t, tt.wantResult, result)
		})
	}
}

func TestIsContainsIntToInt(t *testing.T) {
	type args struct {
		what  []int
		where []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult bool
	}{
		{
			name:       "arrayIntInArrayInt",
			args:       args{what: []int{3, 4, 5}, where: []int{13, 14, 3, 4, 5}},
			wantResult: true,
		},
		{
			name:       "arrayIntInArrayIntFalse",
			args:       args{what: []int{3, 4, 6}, where: []int{13, 14, 3, 4, 5}},
			wantResult: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isContainsIntToInt(tt.args.what, tt.args.where)
			assert.Equal(t, tt.wantResult, result)
		})
	}
}
