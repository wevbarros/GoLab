package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	arr := make([]int, 100)
	for i := range arr {
		arr[i] = rand.Intn(1000)
	}
	fmt.Println("Array original:", arr)

	chunkSize := len(arr) / 2

	arr1 := arr[:chunkSize]
	arr2 := arr[chunkSize:]

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		sort.Ints(arr1)
		wg.Done()
	}()
	go func() {
		sort.Ints(arr2)
		wg.Done()
	}()

	wg.Wait()

	result := merge(arr1, arr2)
	fmt.Println("Array ordenado:", result)
}

func merge(arr1, arr2 []int) []int {
	result := make([]int, len(arr1)+len(arr2))
	i, j, k := 0, 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result[k] = arr1[i]
			i++
		} else {
			result[k] = arr2[j]
			j++
		}
		k++
	}
	for i < len(arr1) {
		result[k] = arr1[i]
		i++
		k++
	}
	for j < len(arr2) {
		result[k] = arr2[j]
		j++
		k++
	}
	return result
}
