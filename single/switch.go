package single

func SwitchType(what interface{}, where interface{}) bool {
	switch what.(type) {
	case int:
		if _, ok := where.([]int); ok {
			return isContainsIntToArrayInt(what.(int), where.([]int))
		}
		return tryConvertForInt(what, where)
	case string:
		if _, ok := where.([]string); ok {
			return isContainsStringToArrayString(what.(string), where.([]string))
		}
		return tryConvertForString(what, where)
	default:
		return false
	}
}
