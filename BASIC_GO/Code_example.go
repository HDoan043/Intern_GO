package main

import (
	"fmt"
	"time"
	// "error"
)

// Global variables
var length int = 10000

// Define new struct
type NewStruct struct {
	Id   int
	Name string
}

// Use Receiver to link the function to the struct
func (ns *NewStruct) ChangeStruct(newId int, newName string) {
	ns.Id = newId
	ns.Name = newName
}

// Global function
func SingleWorker(arr []int) int {
	sum := 0
	for i := range arr {
		sum += i
		// Waiting for 0.25 seconds to similate an IO task
		time.Sleep(1)
	}

	return sum
}

func main() {
	// Array
	arr1 := [...]int{1, 2, 3}
	arr2 := arr1
	arr2[0] = 100
	fmt.Printf("arr1 has %d elements: %v\n", len(arr1), arr1)
	fmt.Printf("arr2 has %d elements: %v\n", len(arr2), arr2)
	fmt.Println()

	// Struct
	object1 := NewStruct{
		Id:   1,
		Name: "A",
	}
	fmt.Printf("object1 before change: %v\n", object1)
	object2 := object1
	fmt.Printf("object2 before change: %+v\n", object2)

	object2.ChangeStruct(2, "B")
	fmt.Printf("object1 after change: %#v\n", object1)
	fmt.Printf("object2 after change: %#v\n", object2)
	fmt.Println()

	// Slice
	var slice1 = []int{1, 2, 3, 4}
	slice2 := make([]int, 4, 5)
	slice3 := slice1
	slice4 := slice2
	fmt.Printf("slice1 before change: len: %d | cap: %d| slice1: %v\n", len(slice1), cap(slice1), slice1)
	fmt.Println("slice2 before change: len:", len(slice2), "| cap:", cap(slice2), "| slice2: ", slice2)
	fmt.Println("slice3 before change: len:", len(slice3), "| cap:", cap(slice3), "| slice3: ", slice3)
	fmt.Println("slice4 before change: len:", len(slice4), "| cap:", cap(slice4), "| slice4: ", slice4)
	slice3[0] = 100
	slice4 = append(slice4, 6)
	fmt.Printf("slice1 before change: len: %d | cap: %d| slice1: %v\n", len(slice1), cap(slice1), slice1)
	fmt.Println("slice2 before change: len:", len(slice2), "| cap:", cap(slice2), "| slice2: ", slice2)
	fmt.Println("slice3 before change: len:", len(slice3), "| cap:", cap(slice3), "| slice3: ", slice3)
	fmt.Println("slice4 before appending 1 element: len:", len(slice4), "| cap:", cap(slice4), "| slice4: ", slice4)
	slice4 = append(slice4, 7)
	fmt.Println("slice4 before appending 2 elemetns: len:", len(slice4), "| cap:", cap(slice4), "| slice4: ", slice4)
	fmt.Println()

	// Map
	map1 := map[string]int{
		"A": 1,
		"B": 2,
	}
	map2 := map1
	fmt.Println("map1 before change:", map1)
	fmt.Println("map2 before change:", map2)
	map2["C"] = 3
	delete(map1, "A")
	_, okB := map1["B"]
	_, okA := map1["A"]
	fmt.Println("map1 after change: ", map1)
	fmt.Println("map2 after change: ", map2)
	fmt.Println("map1 after change has key A: ", okA)
	fmt.Println("map1 after change has key B: ", okB)
	fmt.Println()

	// Local function and anonymous function
	localFunc := func(a int, b int) int {
		fmt.Printf("This is local function: receive %d and %d, ", a, b)
		return a + b
	}
	resLocalFunc := localFunc(1, 2)
	fmt.Println("returns ", resLocalFunc)

	resAnomFunc := func(a int, b int) int {
		fmt.Printf("This is anonymous function: receive %d and %d, ", a, b)
		return a - b
	}(3, 5)
	fmt.Println("return ", resAnomFunc)

	// Concurrence Processing
	tasks := map[string][]int{
		"task1": {1, 2},
		"task2": {3, 4, 5},
		"task3": {6, 7, 8, 9},
	}

	// Sequential Processing
	sum := 0
	start1 := time.Now()
	for _, taskValue := range tasks {
		sum += SingleWorker(taskValue)
	}
	end1 := time.Since(start1)
	fmt.Println("Sequential processing returns sum of the slice being", sum, "in", end1)

	// Concurrence Processing
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	chList := []chan int{ch1, ch2, ch3}
	sum = 0

	singleTask := func(arr []int, ch chan int) {
		ch <- SingleWorker(arr)
	}

	length := len(tasks)
	taskValueList := make([][]int, 0, length)
	for _, taskValue := range tasks {
		taskValueList = append(taskValueList, taskValue)
	}

	start2 := time.Now()
	for i := range length {
		go singleTask(taskValueList[i], chList[i])
	}

	for _, ch := range chList {
		sum += <-ch
	}
	end2 := time.Since(start2)
	fmt.Println("Concurrence processing returns sum of the slice being", sum, "in", end2)
}
