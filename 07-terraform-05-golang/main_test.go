package main

import (
	"fmt"
	"testing"
)


// Тест задачи 1
func TestTask1(t *testing.T) {
	// Тесттовые данные
	var meter float64 = 2
	var v float64 = task1(meter)
	// Возвращаемое значение должно быть > 0
	if v < 0 {
		t.Error("Expected > 0, got ", v)
	}
}


// Тест задачи 2
func TestTask2(t *testing.T) {
	// Исходный массив для теста (ожидаем значение > 0)
	var x = []int{48,96,86,68,57,82,63,70,37,34,83,27,19,97,9,17}
	_, v := task2(x)
	// Ожидаем значение > 0
	if v < 0 {
		t.Error("Expected >= 0, got ", v)
	}

	// Исходный массив для теста (ожидаем значение < 0)
	var y = []int{48,96,86,68,57,82,63,70,37,34,83,27,19,97,9,17,0,-2}
	_, v2 := task2(y)
	// Ожидаем значение < 0
	if v2 > 0 {
		t.Error("Expected < 0, got ", v2)
	}
}


// Тест задачи 3
func TestTask3(t *testing.T) {
	// Число, на которое делится последовательность
	var x int = 4
	_, arr := task3(x)
	// Первый элемент массива должен быть равен числу, на которое делим последовательность
	if arr[0] != x {
		var xi string = fmt.Sprintf("%d",x)
		t.Error("Expected " + xi + ", got ", arr[0])
	}
}
