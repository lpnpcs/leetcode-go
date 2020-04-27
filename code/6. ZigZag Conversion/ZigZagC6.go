package src

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	rows := make([]string, numRows)
	index := 0
	dir := 1
	for _, val := range s {
		rows[index] += string(val)
		index += dir
		if index == 0 || index == numRows-1 {
			dir = -dir
		}

	}
	var res string
	for _, val := range rows {
		res += val
	}
	return res

}
