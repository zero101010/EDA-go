package main

func search(nums []int, target int) int {
	// inicio do array
	baixo := 0
	// final do array
	alto := len(nums) - 1
	for baixo <= alto {
		// valor do meio é 4
		meio := alto + baixo
		chute := nums[meio]
		if target == chute {
			return nums[meio]
		}
		if target < chute {
			alto = meio - 1
		} else {
			baixo = meio + 1
		}
	}
	return -1

}
