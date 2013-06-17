package array

func InArray(need string, itemList []string) bool {
	result := false
	for _, item := range itemList {
		if item == need {
			result = true
		}
	}
	return result
}

func ArrayKeyExists() {

}
