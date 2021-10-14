package single

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
			name:       "singleIntInArrayString",
			args:       args{what: 3, where: []string{"13", "14", "3"}},
			wantResult: true,
		},
		{
			name:       "singleIntInArrayStringFalse",
			args:       args{what: 3, where: []string{"13", "14", "4"}},
			wantResult: false,
		},
		{
			name:       "singleIntInArrayFloat32",
			args:       args{what: 3, where: []float32{13, 14, 3}},
			wantResult: true,
		},
		{
			name:       "singleIntInArrayFloat32False",
			args:       args{what: 3, where: []float32{13, 14, 3.33}},
			wantResult: false,
		},
		{
			name:       "singleIntInArrayFloat64",
			args:       args{what: 3, where: []float64{13, 14, 3}},
			wantResult: true,
		},
		{
			name:       "singleIntInArrayFloat64False",
			args:       args{what: 3, where: []float64{13, 14, 3.1}},
			wantResult: false,
		},
		{
			name:       "singleIntInSingleString",
			args:       args{what: 3, where: "number 3 - the third"},
			wantResult: true,
		},
		{
			name:       "singleIntInSingleStringFalse",
			args:       args{what: 10, where: "number 3 - the third"},
			wantResult: false,
		},
		{
			name:       "singleIntInArrayUint",
			args:       args{what: 1, where: []uint{1, 2, 3}},
			wantResult: true,
		},
		{
			name:       "singleIntInArrayUintFalse",
			args:       args{what: 5, where: []uint{1, 2, 3}},
			wantResult: false,
		},
		{
			name:       "singleIntInArrayUint8",
			args:       args{what: 1, where: []uint8{1, 2, 3}},
			wantResult: true,
		},
		{
			name:       "singleIntInArrayUint8False",
			args:       args{what: 5, where: []uint8{1, 2, 3}},
			wantResult: false,
		},
		{
			name:       "singleIntInArrayUint16",
			args:       args{what: 1, where: []uint16{1, 2, 3}},
			wantResult: true,
		},
		{
			name:       "singleIntInArrayUint16False",
			args:       args{what: 5, where: []uint16{1, 2, 3}},
			wantResult: false,
		},
		{
			name:       "singleIntInArrayUint32",
			args:       args{what: 1, where: []uint32{1, 2, 3}},
			wantResult: true,
		},
		{
			name:       "singleIntInArrayUint32False",
			args:       args{what: 5, where: []uint32{1, 2, 3}},
			wantResult: false,
		},
		{
			name:       "singleIntInArrayUint64",
			args:       args{what: 1, where: []uint64{1, 2, 3}},
			wantResult: true,
		},
		{
			name:       "singleIntInArrayUint64False",
			args:       args{what: 5, where: []uint64{1, 2, 3}},
			wantResult: false,
		},
		{
			name:       "singleIntInSingleInt",
			args:       args{what: 3, where: 3},
			wantResult: true,
		},
		{
			name:       "singleIntInSingleIntFalse",
			args:       args{what: 3, where: 4},
			wantResult: false,
		},
		{
			name:       "singleIntInSingleString",
			args:       args{what: 3, where: "333"},
			wantResult: true,
		},
		{
			name:       "singleIntInSingleStringFalse",
			args:       args{what: 3, where: "4"},
			wantResult: false,
		},
		{
			name:       "singleIntInSingleFloat32",
			args:       args{what: 3, where: float32(3.0)},
			wantResult: true,
		},
		{
			name:       "singleIntInSingleFloat32False",
			args:       args{what: 3, where: float32(4)},
			wantResult: false,
		},
		{
			name:       "singleIntInSingleFloat64",
			args:       args{what: 3, where: float64(3)},
			wantResult: true,
		},
		{
			name:       "singleIntInSingleFloat64False",
			args:       args{what: 3, where: float64(4)},
			wantResult: false,
		},
		{
			name:       "singleIntInSingleUint",
			args:       args{what: 3, where: uint(3)},
			wantResult: true,
		},
		{
			name:       "singleIntInSingleUintFalse",
			args:       args{what: 3, where: uint(4)},
			wantResult: false,
		},
		{
			name:       "singleIntInSingleUint8",
			args:       args{what: 3, where: uint8(3)},
			wantResult: true,
		},
		{
			name:       "singleIntInSingleUint8False",
			args:       args{what: 3, where: uint8(4)},
			wantResult: false,
		},
		{
			name:       "singleIntInSingleUint16",
			args:       args{what: 3, where: uint16(3)},
			wantResult: true,
		},
		{
			name:       "singleIntInSingleUint16False",
			args:       args{what: 3, where: uint16(4)},
			wantResult: false,
		},
		{
			name:       "singleIntInSingleUint32",
			args:       args{what: 3, where: uint32(3)},
			wantResult: true,
		},
		{
			name:       "singleIntInSingleUint32False",
			args:       args{what: 3, where: uint32(4)},
			wantResult: false,
		},
		{
			name:       "singleIntInSingleUint64",
			args:       args{what: 3, where: uint64(3)},
			wantResult: true,
		},
		{
			name:       "singleIntInSingleUint64False",
			args:       args{what: 3, where: uint64(4)},
			wantResult: false,
		},
		{
			name:       "singleIntInUndefinedType",
			args:       args{what: 3, where: func() {}},
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

func TestIsContainsIntToArrayInt(t *testing.T) {
	type args struct {
		what  int
		where []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult bool
	}{
		{
			name:       "singleIntInArrayInt",
			args:       args{what: 3, where: []int{13, 14, 3}},
			wantResult: true,
		},
		{
			name:       "singleIntInArrayIntFalse",
			args:       args{what: 3, where: []int{13, 14, 33}},
			wantResult: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isContainsIntToArrayInt(tt.args.what, tt.args.where)
			assert.Equal(t, tt.wantResult, result)
		})
	}
}
