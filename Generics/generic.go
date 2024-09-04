package main

import "fmt"

type Number interface{
	int | int16 | int32 | int64 | float32 | float64
}

func main() {
	x := Sum(2, 7)
	fmt.Printf("%d\n", x)

	y := Sum(2.5, 7.5)
	fmt.Printf("%f\n", y)
}

func SumInts(a, b int) int {
	return a + b
}

func SumFloats(a, b float64) float64 {
	return a + b
}

func Sum[T Number](a, b T) T {
	return a + b
}



func sliceSum() {

	mySlice := []int{1, 2, 3, 4, 5}

	fmt.Printf("%d\n", SumSlice(mySlice))

	mySlice2 := []float64{1.5, 2.5, 3.5, 4.5, 5.5}

	fmt.Printf("%f\n", SumSlice(mySlice2))
}

func SumSlice[T Number](slc []T) (sum T) {
	for _, v := range slc {
		sum += v
	}
	return
}
