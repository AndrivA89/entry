package array

import "strconv"

func isContainsIntToInt(what []int, where []int) bool {
	ok := make([]bool, 0, len(what))
	for _, o := range what {
		for _, w := range where {
			if o == w {
				ok = append(ok, true)
				break
			}
		}
	}
	if len(ok) < len(what) {
		return false
	}
	return true
}

func isContainsIntToString(what []int, where []string) bool {
	ok := make([]bool, 0, len(what))
	for _, o := range what {
		for _, w := range where {
			if strconv.Itoa(o) == w {
				ok = append(ok, true)
				break
			}
		}
	}
	if len(ok) < len(what) {
		return false
	}
	return true
}

func isContainsIntToFloat64(what []int, where []float64) bool {
	ok := make([]bool, 0, len(what))
	for _, o := range what {
		for _, w := range where {
			if float64(o) == w {
				ok = append(ok, true)
				break
			}
		}
	}
	if len(ok) < len(what) {
		return false
	}
	return true
}

func isContainsIntToFloat32(what []int, where []float32) bool {
	ok := make([]bool, 0, len(what))
	for _, o := range what {
		for _, w := range where {
			if float32(o) == w {
				ok = append(ok, true)
				break
			}
		}
	}
	if len(ok) < len(what) {
		return false
	}
	return true
}

func isContainsIntToUint(what []int, where []uint) bool {
	ok := make([]bool, 0, len(what))
	for _, o := range what {
		for _, w := range where {
			if uint(o) == w {
				ok = append(ok, true)
				break
			}
		}
	}
	if len(ok) < len(what) {
		return false
	}
	return true
}

func isContainsIntToUint8(what []int, where []uint8) bool {
	ok := make([]bool, 0, len(what))
	for _, o := range what {
		for _, w := range where {
			if uint8(o) == w {
				ok = append(ok, true)
				break
			}
		}
	}
	if len(ok) < len(what) {
		return false
	}
	return true
}

func isContainsIntToUint16(what []int, where []uint16) bool {
	ok := make([]bool, 0, len(what))
	for _, o := range what {
		for _, w := range where {
			if uint16(o) == w {
				ok = append(ok, true)
				break
			}
		}
	}
	if len(ok) < len(what) {
		return false
	}
	return true
}

func isContainsIntToUint32(what []int, where []uint32) bool {
	ok := make([]bool, 0, len(what))
	for _, o := range what {
		for _, w := range where {
			if uint32(o) == w {
				ok = append(ok, true)
				break
			}
		}
	}
	if len(ok) < len(what) {
		return false
	}
	return true
}

func isContainsIntToUint64(what []int, where []uint64) bool {
	ok := make([]bool, 0, len(what))
	for _, o := range what {
		for _, w := range where {
			if uint64(o) == w {
				ok = append(ok, true)
				break
			}
		}
	}
	if len(ok) < len(what) {
		return false
	}
	return true
}

func tryConvertForInt(object interface{}, where interface{}) bool {
	switch where.(type) {
	case []string:
		return isContainsIntToString(object.([]int), where.([]string))
	case []float64:
		return isContainsIntToFloat64(object.([]int), where.([]float64))
	case []float32:
		return isContainsIntToFloat32(object.([]int), where.([]float32))
	case []uint:
		return isContainsIntToUint(object.([]int), where.([]uint))
	case []uint8:
		return isContainsIntToUint8(object.([]int), where.([]uint8))
	case []uint16:
		return isContainsIntToUint16(object.([]int), where.([]uint16))
	case []uint32:
		return isContainsIntToUint32(object.([]int), where.([]uint32))
	case []uint64:
		return isContainsIntToUint64(object.([]int), where.([]uint64))
	default:
		return false
	}
}
