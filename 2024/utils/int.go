package utils

import "strconv"

func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

func IntAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ToIntSlice(s []string) []int {
	var arr []int
	for _, v := range s {
		arr = append(arr, ToInt(v))
	}
	return arr
}
