package array

func SwitchType(what interface{}, where interface{}) bool {
	switch what.(type) {
	case []int:
		if _, ok := where.([]int); ok {
			return isContainsIntToInt(what.([]int), where.([]int))
		}
		return tryConvertForInt(what, where)
	case []string:
		return false
	default:
		return false
	}
}
