package main

import "fmt"

var x = []int{48,96,86,68,57,82,63,70,37,34,83,27,19,97,9,17}


// Задача 1 (метры в футы)
func task1(meter float64) float64 {
	return meter / 0.3048
}


// Задача 2 (наименьшее значение)
func task2(x []int) (string, int) {
	var min int = x[0]
	for _, v := range x {
		if v < min {
			min = v
		}
	}
	return "Наименьшее значение:", min
}


// Задача 3 (числа от 1 до 100, которые делятся на ...)
func task3(d int) (string, []int) {
	var arr []int
	for i := 1; i<=100; i++ {
		if i % d == 0 {
			arr = append(arr, i)
		}
	}
	var di string = fmt.Sprintf("%d",d)

	return "числа от 1 до 100, которые делятся на " + di + ":", arr
}


func main() {
	// Задача 1
	fmt.Print("Введите длину в метрах: ")
	var meter float64
	fmt.Scanf("%f", &meter)

	t1res := task1(meter)
	fmt.Println("Задание 1 (метры в футы)")
	fmt.Println(t1res)

	// Задача 2
	t2res1, t2res2 := task2(x)
	fmt.Println("Задание 2 (наименьшее значение)")
	fmt.Println(t2res1, t2res2)

	// Задача 3
	t3res1, t3res2 := task3(3)
	fmt.Println("Задание 3 (числа от 1 до 100, которые делятся на ...)")
	fmt.Println(t3res1, t3res2)
}
