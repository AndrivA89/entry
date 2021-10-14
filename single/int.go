package single

import (
	"strconv"
	"strings"
)

func isContainsIntToArrayInt(what int, where []int) bool {
	for _, v := range where {
		if what == v {
			return true
		}
	}
	return false
}

func isContainsIntToArrayString(what int, where []string) bool {
	for _, v := range where {
		if strconv.Itoa(what) == v {
			return true
		}
	}
	return false
}

func isContainsIntToArrayFloat64(what int, where []float64) bool {
	for _, v := range where {
		if float64(what) == v {
			return true
		}
	}
	return false
}

func isContainsIntToArrayFloat32(what int, where []float32) bool {
	for _, v := range where {
		if float32(what) == v {
			return true
		}
	}
	return false
}

func isContainsIntToArrayUint(what int, where []uint) bool {
	for _, v := range where {
		if uint(what) == v {
			return true
		}
	}
	return false
}

func isContainsIntToArrayUint8(what int, where []uint8) bool {
	for _, v := range where {
		if uint8(what) == v {
			return true
		}
	}
	return false
}

func isContainsIntToArrayUint16(what int, where []uint16) bool {
	for _, v := range where {
		if uint16(what) == v {
			return true
		}
	}
	return false
}

func isContainsIntToArrayUint32(what int, where []uint32) bool {
	for _, v := range where {
		if uint32(what) == v {
			return true
		}
	}
	return false
}

func isContainsIntToArrayUint64(what int, where []uint64) bool {
	for _, v := range where {
		if uint64(what) == v {
			return true
		}
	}
	return false
}

func tryConvertForInt(what interface{}, where interface{}) bool {
	switch where.(type) {
	case []string:
		return isContainsIntToArrayString(what.(int), where.([]string))
	case []float64:
		return isContainsIntToArrayFloat64(what.(int), where.([]float64))
	case []float32:
		return isContainsIntToArrayFloat32(what.(int), where.([]float32))
	case []uint:
		return isContainsIntToArrayUint(what.(int), where.([]uint))
	case []uint8:
		return isContainsIntToArrayUint8(what.(int), where.([]uint8))
	case []uint16:
		return isContainsIntToArrayUint16(what.(int), where.([]uint16))
	case []uint32:
		return isContainsIntToArrayUint32(what.(int), where.([]uint32))
	case []uint64:
		return isContainsIntToArrayUint64(what.(int), where.([]uint64))
	case int:
		return where.(int) == what.(int)
	case string:
		return strings.Contains(strings.ToLower(where.(string)), strconv.Itoa(what.(int)))
	case float64:
		return float64(what.(int)) == where.(float64)
	case float32:
		return float32(what.(int)) == where.(float32)
	case uint:
		return uint(what.(int)) == where.(uint)
	case uint8:
		return uint8(what.(int)) == where.(uint8)
	case uint16:
		return uint16(what.(int)) == where.(uint16)
	case uint32:
		return uint32(what.(int)) == where.(uint32)
	case uint64:
		return uint64(what.(int)) == where.(uint64)
	default:
		return false
	}
}
