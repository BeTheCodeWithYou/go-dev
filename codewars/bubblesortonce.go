package main

func bubbleSortOnce(i []int) []int {
	num := make([]int, len(i))
	copy(num, i)
	for k := 0; k < len(num)-1; k++ {
		if num[k] > num[k+1] {
			num[k], num[k+1] = num[k+1], num[k]
		}
	}
	return num
}
