package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	fmt.Println("My favourite answer is", rand.Intn(10))

	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}

	fmt.Println(add(3, 4))
	fmt.Println(sqrt(10.0))
	fmt.Println(pow(2, 3, 10))
	fmt.Println((newton(23)))
	checkOs()
	switchelse()
}

func add(x, y int) int {
	return x + y
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))

}

func pow(x, n, lim float64) float64 {
	//note the use of v here, called a short statement
	//only avaiable in if and else block\

	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >=%g\n", v, lim)
	}

	return lim
}

func newton(x float64) float64 {
	z := 1.0

	if math.Pow(z, 2) == x {
		return z
	}

	for math.Abs(z-math.Sqrt(x)) > 0.000001 {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func checkOs() {

	switch os := runtime.GOOS; os {
	case "darwin":
		myos := "OS X"
		fmt.Println("GO is currently running on", myos)
	case "linux":
		myos := "Linux"
		fmt.Println("GO is currently running on", myos)
	}

}

//using empty switch instead of if&else statements
func switchelse() {
	t := time.Now()

	switch{
	case t.Hour() < 12:
		fmt.Println("GM")
	case t.Hour() < 17:
		fmt.Println("GA")
	default:
		fmt.Println("Good Evening")
	}

}