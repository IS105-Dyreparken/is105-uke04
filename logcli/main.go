package main

import (
	"os"
  "fmt"
  "strconv"
  "math"
)

func main(){
  arg1 := os.Args[1]
  arg1Float, _ := strconv.ParseFloat(arg1, 64)
  fmt.Println(math.Log2(arg1Float))
  }
