package single

import (
	"fmt"
	"strconv"
	"strings"
)

func isContainsStringToArrayString(what string, where []string) bool {
	for _, v := range where {
		if what == v {
			return true
		}
	}
	return false
}

func isContainsStringToArrayInt(what string, where []int) bool {
	for _, v := range where {
		if what == strconv.Itoa(v) {
			return true
		}
	}
	return false
}

func isContainsStringToArrayFloat32(what string, where []float32) bool {
	for _, v := range where {
		if what == fmt.Sprintf("%v", v) {
			return true
		}
	}
	return false
}

func isContainsStringToArrayFloat64(what string, where []float64) bool {
	for _, v := range where {
		if what == fmt.Sprintf("%v", v) {
			return true
		}
	}
	return false
}

func isContainsStringToArrayUint(what string, where []uint) bool {
	for _, v := range where {
		if what == fmt.Sprintf("%v", v) {
			return true
		}
	}
	return false
}

func isContainsStringToArrayUint8(what string, where []uint8) bool {
	for _, v := range where {
		if what == fmt.Sprintf("%v", v) {
			return true
		}
	}
	return false
}

func isContainsStringToArrayUint16(what string, where []uint16) bool {
	for _, v := range where {
		if what == fmt.Sprintf("%v", v) {
			return true
		}
	}
	return false
}

func isContainsStringToArrayUint32(what string, where []uint32) bool {
	for _, v := range where {
		if what == fmt.Sprintf("%v", v) {
			return true
		}
	}
	return false
}

func isContainsStringToArrayUint64(what string, where []uint64) bool {
	for _, v := range where {
		if what == fmt.Sprintf("%v", v) {
			return true
		}
	}
	return false
}

func tryConvertForString(what interface{}, where interface{}) bool {
	switch where.(type) {
	case []int:
		return isContainsStringToArrayInt(what.(string), where.([]int))
	case []float32:
		return isContainsStringToArrayFloat32(what.(string), where.([]float32))
	case []float64:
		return isContainsStringToArrayFloat64(what.(string), where.([]float64))
	case []uint:
		return isContainsStringToArrayUint(what.(string), where.([]uint))
	case []uint8:
		return isContainsStringToArrayUint8(what.(string), where.([]uint8))
	case []uint16:
		return isContainsStringToArrayUint16(what.(string), where.([]uint16))
	case []uint32:
		return isContainsStringToArrayUint32(what.(string), where.([]uint32))
	case []uint64:
		return isContainsStringToArrayUint64(what.(string), where.([]uint64))
	case string:
		return strings.Contains(strings.ToLower(where.(string)), strings.ToLower(what.(string)))
	case int:
		return strings.ToLower(fmt.Sprintf("%v", where.(int))) == strings.ToLower(what.(string))
	case float32:
		return strings.ToLower(fmt.Sprintf("%v", where.(float32))) == strings.ToLower(what.(string))
	case float64:
		return strings.ToLower(fmt.Sprintf("%v", where.(float64))) == strings.ToLower(what.(string))
	case uint:
		return strings.ToLower(fmt.Sprintf("%v", where.(uint))) == strings.ToLower(what.(string))
	case uint8:
		return strings.ToLower(fmt.Sprintf("%v", where.(uint8))) == strings.ToLower(what.(string))
	case uint16:
		return strings.ToLower(fmt.Sprintf("%v", where.(uint16))) == strings.ToLower(what.(string))
	case uint32:
		return strings.ToLower(fmt.Sprintf("%v", where.(uint32))) == strings.ToLower(what.(string))
	case uint64:
		return strings.ToLower(fmt.Sprintf("%v", where.(uint64))) == strings.ToLower(what.(string))
	default:
		return false
	}
}
