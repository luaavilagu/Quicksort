package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var count int

func quickSort(arr1 []int) []int {
	if len(arr1) == 1 {
		return arr1
	}

	// if len(arr1) == 2 {
	// 	arr2 := make([]int, 2)
	// 	// count++
	// 	if arr1[0] > arr1[1] {
	// 		arr2[0], arr2[1] = arr1[1], arr1[0]
	// 		return arr2
	// 	}
	// 	return arr1
	// }

	idx := rand.Intn(len(arr1))
	pivote := arr1[idx]
	leftArr := make([]int, 0)
	rightArr := make([]int, 0)
	for _, v := range append(arr1[:idx], arr1[idx+1:]...) {
		if pivote > v {
			leftArr = append(leftArr, v)
		} else {
			rightArr = append(rightArr, v)
		}

	}
	count += len(arr1) - 1
	if len(leftArr) > 0 {
		leftArr = quickSort(leftArr)
	}

	if len(rightArr) > 0 {
		rightArr = quickSort(rightArr)
	}

	return append((append(leftArr, pivote)), rightArr...)

}

func createArray(n int) []int {
	newArr := make([]int, 0)
	for i := 0; i < n; i++ {
		randNum := rand.Intn(1000)
		for contains(newArr, randNum) {
			randNum = rand.Intn(1000)
		}
		newArr = append(newArr, randNum)
	}
	return newArr
}

func contains(arr []int, randInt int) bool {
	for _, a := range arr {
		if a == randInt {
			return true
		}
	}
	return false
}

func main() {

	//arr1 := []int{5, 9, 3, 10, 11, 14, 8, 4, 17, 6}
	//arr1 := []int{4, 3, 2, 1}
	rand.Seed(int64(time.Now().UnixNano()))
	var evaluations, acum float64
	fmt.Println("tam\tpromedio\tesperado\tworst-case")
	for tam := 1; tam < 15; tam++ {
		evaluations = 100000
		for i := 0; i < int(evaluations); i++ {
			count = 0
			arr1 := createArray(tam)
			//fmt.Println(arr1)
			quickSort(arr1)
			acum += float64(count)

		}

		x := math.Log(float64(tam))
		y := math.Log(2.0)

		esperado := 2 * float64(tam) * float64(x/y)
		worstCase := tam * tam
		fmt.Printf("%3d\t%4.2f\t\t%4.2f\t\t%3d\n", tam, acum/evaluations, esperado, worstCase)
		//fmt.Println(acum / evaluations)
	}

	count = 0
	arr1 := []int{5, 9, 3, 10, 11, 14, 8, 4, 17, 6}
	h := quickSort(arr1)
	fmt.Println(h, count)

}
