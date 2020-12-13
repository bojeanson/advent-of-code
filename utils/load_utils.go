package utils

func LoadDataFromFileAndSortThem(filename string) []int {
	var lines []int
	for inputInt := range IntIteratorFromFile(filename) {
		switch {
		case lines == nil || inputInt > lines[len(lines)-1]:
			lines = append(lines, inputInt)
		default:
			for i, _ := range lines {
				if inputInt <= lines[i] {
					lines = InsertInSliceAtPosition(lines, i, inputInt)
					break
				}
			}
		}
	}
	return lines
}

func LoadDataFromFile(filename string) []string {
	var lines []string
	for inputString := range StringIteratorFromFile(filename) {
		lines = append(lines, inputString)
	}
	return lines
}
