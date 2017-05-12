package functions

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func LogaritmeBase2() {
	tall := os.Args[1]
	tallFloat, _ := strconv.ParseFloat(tall, 64)
	fmt.Println("Logaritmen av ", tall, " blir ", math.Log2(tallFloat), " med base 2")
}
