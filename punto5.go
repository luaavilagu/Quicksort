package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var count int

//swap  intercambia elementos de un arreglo
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

//partition particiona el arreglo dejando a los menores
//del pivote adelante y los mayores atras
func partition(arr []int, low, hight int) int {
	pivot := arr[hight]
	i := low - 1
	for j := low; j < hight-1; j++ {
		count++
		if arr[j] <= pivot {
			i++
			swap(arr, i, j)

		}
	}
	swap(arr, hight, i+1)
	return i + 1
}

//partitionR Escoge el pivote aleatoriamente
func partitionR(arr []int, low, hight int) int {
	indexP := rand.Intn(hight)
	swap(arr, indexP, hight)
	return partition(arr, low, hight)
}

//quickSort2 implementado de forma diferente, consume menos memoria
func quickSort2(arr []int, low, hight int) {
	if low < hight {
		p := partitionR(arr, low, hight)

		quickSort2(arr, low, p-1)
		quickSort2(arr, p+1, hight)

	}
}

//quickSort implementado como dice el libro
func quickSort(arr1 []int, low, hight int) []int {
	if len(arr1) == 1 {
		return arr1
	}

	if len(arr1) != 0 {
		idx := rand.Intn(len(arr1))
		pivote := arr1[idx]
		leftArr := make([]int, 0)
		rightArr := make([]int, 0)
		for _, v := range append(arr1[:idx], arr1[idx+1:]...) {
			if v < pivote {
				leftArr = append(leftArr, v)
			} else {
				rightArr = append(rightArr, v)
			}

		}
		count += len(arr1) - 1
		leftArr = quickSort(leftArr, 0, len(leftArr))
		rightArr = quickSort(rightArr, 0, len(leftArr))
		return append((append(leftArr, pivote)), rightArr...)
	}

	return make([]int, 0)

}

// crea un nuevo array sin elementos repetidos
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

//contains verifica si un elemento esta en un arreglo dado
func contains(arr []int, randInt int) bool {
	for _, a := range arr {
		if a == randInt {
			return true
		}
	}
	return false
}

func main() {

	rand.Seed(int64(time.Now().UnixNano()))
	var evaluations, acum float64
	fmt.Println("tam\tpromedio\tesperado\tworst-case")
	for tam := 1; tam < 15; tam++ {
		evaluations = 100000
		for i := 0; i < int(evaluations); i++ {
			count = 0
			arr1 := createArray(tam)
			//fmt.Println(arr1)
			quickSort2(arr1, 0, len(arr1)-1)
			acum += float64(count)

		}

		x := math.Log(float64(tam))
		y := math.Log(2.0)

		esperado := 2 * float64(tam) * float64(x/y)
		worstCase := tam * tam
		fmt.Printf("%3d\t%8.3f\t%8.3f\t%10d\n", tam, acum/evaluations, esperado, worstCase)
		//fmt.Println(acum / evaluations)
	}

	count = 0
	// arr1 := []int{5, 9, 3, 10, 11, 14, 8, 4, 17, 6}
	// h := quickSort2(arr1, 0, len(arr1))
	// fmt.Println(h, count)

}
