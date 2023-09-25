package main

import "fmt"

type City struct {
	name       string
	inhabitant int
	zipCode    string
	surface    int
}

func Form() City {
	var city City
	fmt.Println("Please enter the name of the city:")
	fmt.Scanf("%s\n", &city.name)

	fmt.Println("Please enter the number of inhabitants:")
	fmt.Scanf("%d\n", &city.inhabitant)

	fmt.Println("Please enter the zip code:")
	fmt.Scanf("%s\n", &city.zipCode)

	fmt.Println("Please enter the surface:")
	fmt.Scanf("%d\n", &city.surface)

	return city
}

func main() {
	cities := make(map[string]City)
	var numCities int

	fmt.Println("How many cities do you want to add?")
	fmt.Scanf("%d\n", &numCities)

	for i := 0; i < numCities; i++ {
		city := Form()
		cities[city.name] = city
	}

	fmt.Println("Cities added:")
	for name, city := range cities {
		fmt.Printf("Name: %s, Inhabitants: %d, Zip Code: %s, Surface: %d\n", name, city.inhabitant, city.zipCode, city.surface)
	}
}
