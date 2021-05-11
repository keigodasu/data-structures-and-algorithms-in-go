package main

func FindPair(data []int, value int) bool {
	dic := map[int]bool{}

	for _, v := range data {
		if dic[value-v] {
			return true
		} else {
			dic[v] = true
		}
	}

	return false
}
