package main

import "fmt"

func main() {
	arr := []int{3, 1, 5, 2, 2, 5, 6, 0, 4}
	limit := 2
	// Slice for sum of sub-slices
	var batchSumArr []int

	for j := len(arr) - 2; j >= 0; j-- {
		//Slice the slice:)
		batch := arr[j:min(j+limit, len(arr))]
		//			OR SIMPLY
		//batch := arr[j:j+limit]// j + 2
		
		// How is the code being executed?
		//1.	batch := arr[7:9]
		//2.	batch := arr[6:8]
		//3.	batch := arr[5:7]
		//4.	batch := arr[4:6]
		//5.	batch := arr[3:5]
		//6.	batch := arr[2:4]
		//7.	batch := arr[1:3]
		//8.	batch := arr[0:2]
		fmt.Println(batch)
		batchSum := 0
		for _, v := range batch {
			batchSum += v
		}
		fmt.Println(batchSum)
		batchSumArr = append([]int{batchSum}, batchSumArr...)

	}

	fmt.Println(batchSumArr)
	//Create a   dictionary of values for each element
	dict := make(map[int]int)
	var maxInt int
	for _, num := range batchSumArr {
		dict[num] = dict[num] + 1
	}
	for k, n := range dict {
		if dict[k] > maxInt {
			maxInt = n
		}
	}
	fmt.Println("Dict")
	fmt.Println(dict)
	fmt.Println(maxInt)

}
//Function for slicing slice
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
