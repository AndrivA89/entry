package single

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTryConvertForString(t *testing.T) {
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
			name:       "singleStringInSingleString",
			args:       args{what: "one", where: "One big sentence"},
			wantResult: true,
		},
		{
			name:       "singleStringInArrayInt",
			args:       args{what: "1", where: []int{1, 2, 4}},
			wantResult: true,
		},
		{
			name:       "singleStringInArrayIntFalse",
			args:       args{what: "1", where: []int{11, 12}},
			wantResult: false,
		},
		{
			name:       "singleStringInArrayFloat32",
			args:       args{what: "1.4", where: []float32{1.4, 2, 4}},
			wantResult: true,
		},
		{
			name:       "singleStringInArrayFloat32False",
			args:       args{what: "1", where: []float32{11, 12}},
			wantResult: false,
		},
		{
			name:       "singleStringInArrayFloat64",
			args:       args{what: "16.64", where: []float32{16.64, 2, 4}},
			wantResult: true,
		},
		{
			name:       "singleStringInArrayFloat64PointZero",
			args:       args{what: "1", where: []float64{1.0, 2, 4}},
			wantResult: true,
		},
		{
			name:       "singleStringInArrayFloat64False",
			args:       args{what: "1", where: []float64{1.1, 12}},
			wantResult: false,
		},
		{
			name:       "singleStringInArrayUint",
			args:       args{what: "111", where: []uint{111, 2, 4}},
			wantResult: true,
		},
		{
			name:       "singleStringInArrayUintFalse",
			args:       args{what: "1345", where: []uint{1, 12}},
			wantResult: false,
		},
		{
			name:       "singleStringInArrayUint8",
			args:       args{what: "111", where: []uint8{111, 2, 4}},
			wantResult: true,
		},
		{
			name:       "singleStringInArrayUint8False",
			args:       args{what: "1345", where: []uint8{1, 12}},
			wantResult: false,
		},
		{
			name:       "singleStringInArrayUint16",
			args:       args{what: "111", where: []uint16{111, 2, 4}},
			wantResult: true,
		},
		{
			name:       "singleStringInArrayUint16False",
			args:       args{what: "1345", where: []uint16{1, 12}},
			wantResult: false,
		},
		{
			name:       "singleStringInArrayUint32",
			args:       args{what: "111", where: []uint32{111, 2, 4}},
			wantResult: true,
		},
		{
			name:       "singleStringInArrayUint32False",
			args:       args{what: "1345", where: []uint32{1, 12}},
			wantResult: false,
		},
		{
			name:       "singleStringInArrayUint64",
			args:       args{what: "111", where: []uint64{111, 2, 4}},
			wantResult: true,
		},
		{
			name:       "singleStringInArrayUint64False",
			args:       args{what: "1345", where: []uint64{1, 12}},
			wantResult: false,
		},
		{
			name:       "singleStringInSingleInt",
			args:       args{what: "111", where: 111},
			wantResult: true,
		},
		{
			name:       "singleStringInSingleIntFalse",
			args:       args{what: "1345", where: 111},
			wantResult: false,
		},
		{
			name:       "singleStringInSingleFloat32",
			args:       args{what: "111.1", where: float32(111.1)},
			wantResult: true,
		},
		{
			name:       "singleStringInSingleFloat32False",
			args:       args{what: "1345", where: float32(111.1)},
			wantResult: false,
		},
		{
			name:       "singleStringInSingleFloat64",
			args:       args{what: "111.1", where: 111.1},
			wantResult: true,
		},
		{
			name:       "singleStringInSingleFloat64False",
			args:       args{what: "1345", where: 111.1},
			wantResult: false,
		},
		{
			name:       "singleStringInSingleUint",
			args:       args{what: "1", where: uint(1)},
			wantResult: true,
		},
		{
			name:       "singleStringInSingleUintFalse",
			args:       args{what: "1345", where: uint(1)},
			wantResult: false,
		},
		{
			name:       "singleStringInSingleUint8",
			args:       args{what: "1", where: uint8(1)},
			wantResult: true,
		},
		{
			name:       "singleStringInSingleUint8False",
			args:       args{what: "1345", where: uint8(1)},
			wantResult: false,
		},
		{
			name:       "singleStringInSingleUint16",
			args:       args{what: "1", where: uint16(1)},
			wantResult: true,
		},
		{
			name:       "singleStringInSingleUint16False",
			args:       args{what: "1345", where: uint16(1)},
			wantResult: false,
		},
		{
			name:       "singleStringInSingleUint32",
			args:       args{what: "1", where: uint32(1)},
			wantResult: true,
		},
		{
			name:       "singleStringInSingleUint32False",
			args:       args{what: "1345", where: uint32(1)},
			wantResult: false,
		},
		{
			name:       "singleStringInSingleUint64",
			args:       args{what: "1", where: uint64(1)},
			wantResult: true,
		},
		{
			name:       "singleStringInSingleUint64False",
			args:       args{what: "1345", where: uint64(1)},
			wantResult: false,
		},
		{
			name:       "singleStringInUndefinedType",
			args:       args{what: "3", where: func() {}},
			wantResult: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tryConvertForString(tt.args.what, tt.args.where)
			assert.Equal(t, tt.wantResult, result)
		})
	}
}

func TestIsContainsStringToArrayString(t *testing.T) {
	type args struct {
		what  string
		where []string
	}
	tests := []struct {
		name       string
		args       args
		wantResult bool
	}{
		{
			name:       "singleStringInArrayString",
			args:       args{what: "one", where: []string{"one", "two", "six"}},
			wantResult: true,
		},
		{
			name:       "singleStringInArrayStringFalse",
			args:       args{what: "no_one", where: []string{"one", "two", "six"}},
			wantResult: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isContainsStringToArrayString(tt.args.what, tt.args.where)
			assert.Equal(t, tt.wantResult, result)
		})
	}
}
