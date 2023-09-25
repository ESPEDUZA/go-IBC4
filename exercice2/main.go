package main

import "fmt"

func main() {

	var n float64
	var unit rune
	fmt.Println("Enter your number of transactions: ")
	fmt.Scanf("%f", &n) // Corrected to &n and removed \n
	fmt.Println("Time unit: h, m, s: ")
	fmt.Scanf("%c", &unit) // Corrected to %c and &unit and added a space to consume the newline character

	switch unit {
	case 'h':
		fmt.Printf("Volume in 1 second: %.2f transactions \n", n/3600)
		fmt.Printf("Volume in 1 minute: %.2f transactions \n", n/60)

	case 's':
		fmt.Printf("Volume in 1 minute: %.2f transactions \n", n*60)
		fmt.Printf("Volume in 1 hour: %.2f transactions \n", n*3600)

	case 'm':
		fmt.Printf("Volume in 1 hour: %.2f transactions \n", n*60)
		fmt.Printf("Volume in 1 second: %.2f transactions \n", n/60)

	default:
		fmt.Println("Wrong unit type. Please enter: 'h', 'm' or 's'")
	}
}
