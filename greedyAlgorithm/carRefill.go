package main

import "fmt"

func optimal(capacity int, station []int, n int) int {
	currentStation := 0
	refill := 0
	lastRefill := 0
	for currentStation < n {
		lastRefill = currentStation
		for currentStation < n && (station[currentStation+1]-station[lastRefill]) < capacity {
			currentStation = currentStation + 1

		}

		if currentStation == lastRefill {
			fmt.Println("impossible")
		}
		if currentStation < n {
			refill = refill + 1
		}
	}

	return refill
}

func main() {

	s := []int{0, 200, 600, 750, 1000}
	sonuc := 0
	fmt.Println(sonuc, "sonuç")
	sonuc = optimal(500, s, 7)
	fmt.Println(sonuc, "sonuç")

}
