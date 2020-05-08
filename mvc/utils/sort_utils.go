package utils

func BubbleSort(els []int) []int {

	keepRunning := true

	for keepRunning {

		keepRunning = false

		for i := 0; i < len(els)-1; i++ {

			if els[i] > els[i+1] {
				temp := els[i]
				els[i] = els[i+1]
				els[i+1] = temp
				keepRunning = true
			}
		}
	}

	return els
}

func GetElements(n int) []int {

	els := make([]int, n)

	i := 0
	for j := n-1; j >= 0; j-- {
		els[i] = j
		i++
	}

	return els
}



