package main

import (
	"fmt"
	"log"
	"sort"
	"sync"
)

func main() {
	handleError := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	var wg sync.WaitGroup
	var arraySize int

	semiSortedArray := []int{}

	fmt.Print("How many numbers do you want to add? -> ")
	_, err := fmt.Scan(&arraySize)
	handleError(err)
	array := make([]int, 0, arraySize)

	for i := 0; i < arraySize; i++ {
		var value int
		fmt.Print("Type a value -> ")
		_, err := fmt.Scan(&value)
		handleError(err)
		array = append(array, value)
	}

	fmt.Println("\nArray ->", array)

	subarraysSize := 4

	subarrayLength := len(array) / subarraysSize

	chanel := make(chan []int, int(subarraysSize))

	if len(array) % 2 != 0 {
		semiSortedArray = append(semiSortedArray, array[len(array) - 1])
		fmt.Println("Subarray ->", semiSortedArray)
	}

	for i := 0; i < int(subarraysSize); i++ {
        wg.Add(1)
        go splitAndSortArray(array, subarraysSize, subarrayLength, i, &wg, chanel)
    }

	wg.Wait()
	close(chanel)

	for l := range chanel {
        semiSortedArray = append(semiSortedArray, l...)
    }

    fmt.Println("Semisorted array ->", semiSortedArray)
	sort.Ints(semiSortedArray)
	fmt.Println("Sorted array ->", semiSortedArray)
}

func splitAndSortArray(array []int, subarraysSize int, subarrayLength int, iterator int, wg *sync.WaitGroup, chanel chan []int) {
	arraySplited := array[iterator*subarrayLength : (iterator*subarrayLength)+subarrayLength]
	sort.Ints(arraySplited)
	fmt.Println("Subarray ->", arraySplited)
	chanel <- arraySplited
	wg.Done()
}
