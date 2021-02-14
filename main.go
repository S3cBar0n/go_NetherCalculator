package main

import (
	"fmt"
	"strconv"
)

func netherCoords(coord1 float64, coord3 float64) (float64, float64) {
	coord1 *= 8.0
	coord3 *= 8.0
	return coord1, coord3
}

func overworldCoords(coord1 float64, coord3 float64) (float64, float64) {
	coord1 /= 8.0
	coord3 /= 8.0
	return coord1, coord3
}

func main() {
	var locationStr string
	var location int

	fmt.Print("Enter 1 to convert Nether Coords, Enter 2 to convert Overworld Coords: ")

	for {
		_, err := fmt.Scanln(&locationStr)        // It is setting the entered information above to a string and sees if there is an error printed after the command and sets that to err
		location, err = strconv.Atoi(locationStr) // Converts locationStr to an int and stores it in location, if an error is produced it puts it in err

		if err == nil { // Checks to see if an error was stored in the err variable and if something was, breaks
			if location == 1 || location == 2 {
				var x, y, z float64

				fmt.Print("Enter your x Value: ")
				fmt.Scanf("%f\n", &x)
				fmt.Print("Enter your y Value: ")
				fmt.Scanf("%f\n", &y)
				fmt.Print("Enter your z Value: ")
				fmt.Scanf("%f\n", &z)

				switch location {
				case 1:
					x, z = netherCoords(x, z)
					fmt.Println("Your converted coords are:", "x:", x, "y:", y, "z:", z)
					main()
				case 2:
					x, z = overworldCoords(x, z)
					fmt.Println("Your converted coords are:", "x:", x, "y:", y, "z:", z)
					main()
				default:
					fmt.Println("Please choose 1 or 2, no spaces or other characters allowed...")
					main()
				}
			} else {
				fmt.Println("Please choose 1 or 2, no spaces or other characters allowed...")
				main()
			}
		} else {
			fmt.Println("Please choose 1 or 2, no spaces or other characters allowed...")
			main()
		}
	}
}
