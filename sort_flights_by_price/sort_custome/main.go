// 374.158345ms
package main

import (
	"fmt"
	"sort"
)

// Flight - a struct that
// contains information about flights
type Flight struct {
	Origin      string
	Destination string
	Price       int
}

// function for custom sort ByPrice
// step1. define type
type ByPrice []Flight

// step2. return length
func (b ByPrice) Len() int {
	return len(b)
}

// step3. swap values
func (b ByPrice) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// step4. compare value less return boolen
func (b ByPrice) Less(i, j int) bool {
	return b[i].Price > b[j].Price
}

// SortByPrice sorts flights from highest to lowest
func SortByPrice(flights []Flight) []Flight {
	// implement
	sort.Sort(ByPrice(flights))
	return flights
}

func printFlights(flights []Flight) {
	for _, flight := range flights {
		fmt.Printf("Origin: %s, Destination: %s, Price: %d\n", flight.Origin, flight.Destination, flight.Price)
	}
}

func main() {
	// an slice of flights

	var flights = []Flight{
		Flight{Destination: "Jaipur", Origin: "Pune", Price: 3000},
		Flight{Destination: "Jaipur", Origin: "Mumbai", Price: 3200},
		Flight{Destination: "Jaipur", Origin: "Delhi", Price: 2500},
		Flight{Destination: "Jaipur", Origin: "Cochi", Price: 3500},
		Flight{Destination: "Jaipur", Origin: "Bengaluru", Price: 3750},
	}

	sortedList := SortByPrice(flights)
	printFlights(sortedList)
}
