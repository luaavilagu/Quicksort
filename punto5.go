package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

var count int

//swap  intercambia elementos de un arreglo
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
	count++
}

//partition particiona el arreglo dejando a los menores
//del pivote adelante y los mayores atras
func partition(arr []int, low, hight int) int {
	pivot := arr[hight]
	i := low - 1
	for j := low; j < hight-1; j++ {

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
func quickSort(arr []int, low, hight int) {
	if low < hight {
		p := partitionR(arr, low, hight)

		quickSort(arr, low, p-1)
		quickSort(arr, p+1, hight)

	}
}

//quickSort implementado como dice el libro
// func quickSort(arr1 []int, low, hight int) []int {
// 	if len(arr1) == 1 {
// 		return arr1
// 	}

// 	if len(arr1) != 0 {
// 		idx := rand.Intn(len(arr1))
// 		pivote := arr1[idx]
// 		leftArr := make([]int, 0)
// 		rightArr := make([]int, 0)
// 		for _, v := range append(arr1[:idx], arr1[idx+1:]...) {
// 			if v < pivote {
// 				count++
// 				leftArr = append(leftArr, v)
// 			} else {
// 				rightArr = append(rightArr, v)
// 			}

// 		}

// 		leftArr = quickSort(leftArr, 0, len(leftArr))
// 		rightArr = quickSort(rightArr, 0, len(leftArr))
// 		return append((append(leftArr, pivote)), rightArr...)
// 	}

// 	return make([]int, 0)

// }

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
func createSortedArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return arr
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
	fileName := flag.String("fn", "barchart", "name for file to save barchart image")
	maxSize := flag.Int("size", 10, "max size for the arrays")
	iterations := flag.Int("it", 100000, "number of iterations")
	flag.Parse()
	*fileName += ".png"

	rand.Seed(int64(time.Now().UnixNano()))
	var evaluations, acum float64
	size := *maxSize
	var valorObtenido, valorEsperado, valorPeorCaso plotter.Values
	fmt.Println("tam\tpromedio\tesperado\tworst-case")
	init := 2
	for tam := init; tam <= size; tam++ {
		evaluations = float64(*iterations)
		for i := 0; i < int(evaluations); i++ {
			count = 0
			arr1 := createArray(tam)
			//fmt.Println(arr1)
			quickSort(arr1, 0, len(arr1)-1)
			acum += float64(count)
		}
		x := math.Log10(float64(tam))
		esperado := 2 * float64(tam) * x
		worstCase := tam * tam
		valorObtenido = append(valorObtenido, acum/evaluations)
		valorEsperado = append(valorEsperado, esperado)
		valorPeorCaso = append(valorPeorCaso, float64(worstCase))
		fmt.Printf("%3d\t%8.3f\t%8.3f\t%10d\n", tam, acum/evaluations, esperado, worstCase)
		acum = 0
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "QuickSort"
	p.Y.Label.Text = "Cantidad de Intercambios"
	w := vg.Points(10)

	columnaObtenidos, err := plotter.NewBarChart(valorObtenido, w)
	if err != nil {
		panic(err)
	}
	columnaObtenidos.LineStyle.Width = vg.Length(0)
	columnaObtenidos.Color = plotutil.DarkColors[2]
	columnaObtenidos.XMin = float64(init)
	columnaObtenidos.Offset = -w
	columnaEsperados, err := plotter.NewBarChart(valorEsperado, w)
	if err != nil {
		panic(err)
	}

	columnaEsperados.LineStyle.Width = vg.Length(0)
	columnaEsperados.Color = plotutil.DarkColors[0]
	columnaEsperados.XMin = float64(init)
	columnaPeorCaso, err := plotter.NewBarChart(valorPeorCaso, w)

	if err != nil {
		panic(err)
	}

	columnaPeorCaso.LineStyle.Width = vg.Length(0)
	columnaPeorCaso.Color = plotutil.DarkColors[3]
	columnaPeorCaso.XMin = float64(init)
	columnaPeorCaso.Offset = w

	p.Add(columnaObtenidos, columnaEsperados, columnaPeorCaso)
	p.Legend.Add("Obtenido", columnaObtenidos)
	p.Legend.Add("Esperado", columnaEsperados)
	p.Legend.Add("Peor Caso", columnaPeorCaso)
	p.Legend.Top = true
	p.Legend.Left = true
	p.X.Max = float64(size + 1)
	p.Y.Max = valorPeorCaso[len(valorPeorCaso)-1]

	p.X.Label.Text = "Tamaño del arreglo"
	scale := vg.Length(w * 2)
	if err := p.Save(16*scale*vg.Millimeter, 9*scale*vg.Millimeter, *fileName); err != nil {
		panic(err)
	}
}
