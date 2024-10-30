package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

func main() {
	/*
			var data = []map[string]string{
				{
					"name": "chicken blue", "gender": "male", "color": "red",
				},
				{
					"address": "pegangsaan dua", "number": "25",
				},
				{
					"owner": "logi",
				},
			}

			fmt.Println(data)


		var names = []string{"logi", "tech"}
		printMsg("halo", names)
	*/

	/*
		var randomValue int

		randomValue = randomWithRange(2, 10)
		fmt.Println("random number: ", randomValue)
		randomValue = randomWithRange(2, 10)
		fmt.Println("random number: ", randomValue)
		randomValue = randomWithRange(2, 10)
		fmt.Println("random number: ", randomValue)
	*/

	var diameter float64 = 15

	var area, circumference = calculate(diameter)

	fmt.Printf("luas lingkaran \t\t: %.2f \n", area)
	fmt.Printf("keliling lingkarang\t: %.2f \n", circumference)

}

/*
	func printMsg(msg string, arr []string) {
		var nameString = strings.Join(arr, "")
		fmt.Println(msg, nameString)
	}

	func randomWithRange(max, min int) int {
		var value = randomizer.Int()%(max-min+1) + min
		return value
	}
*/
func calculate(d float64) (float64, float64) {
	var area = math.Pi * math.Pow(d/2, 2)

	var circumference = math.Pi * d

	return area, circumference
}
