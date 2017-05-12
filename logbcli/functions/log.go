package functions

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func LogaritmeFunc() {
	base := os.Args[1]
	tall := os.Args[2]
	tallFloat, _ := strconv.ParseFloat(tall, 64)
	baseFloat, _ := strconv.Atoi(base)

	if baseFloat == 2 {
		fmt.Println("Logaritmen av ", tall, " blir: ", math.Log2(tallFloat), ", beregnet med base 2.")
	} else if baseFloat == 10 {
		fmt.Println("Logaritmen av ",tall, " blir: ", math.Log10(tallFloat), ", beregnet med base 10.")
	} else {
		fmt.Println("Første argumentet må være 2 eller 10")
	}
}
