package main

import "fmt"

// Flight - a struct that
// contains information about flights
type Flight struct {
	Origin      string
	Destination string
	Price       int
}

// SortByPrice sorts flights from highest to lowest
func SortByPrice(flights []Flight) []Flight {
	// implement
	for i := range flights {
		sweeped := false
		for j := 0; j < len(flights)-1-i; j++ {
			if flights[j].Price < flights[j+1].Price {
				flights[j], flights[j+1] = flights[j+1], flights[j]
				sweeped = true
			}
		}
		if !sweeped {
			break
		}
	}
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
