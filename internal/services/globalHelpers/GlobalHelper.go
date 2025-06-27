package globalHelpers

func StringContains(arr []string, target string) int {
	for index, s := range arr {
		if s == target {
			return index
		}
	}
	return -1
}

func IntContains(arr []int, target int) int {
	for index, s := range arr {
		if s == target {
			return index
		}
	}
	return -1
}

func RemoveIntAtIndex(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}
