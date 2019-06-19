package main

import (
	"fmt"
	"strconv"
	"strings"
)

var g = 4

func printGlobal() {
	fmt.Printf("g = %d", g)
}

func printGlobalAndLocal() {
	fmt.Printf("g = %d", g)
	var l = 5
	fmt.Printf("l = %d", l)
}

func foo() *int {
	i := 1
	return &i
}

func usePointerToFunctionLocalVariable() {
	var y *int
	y = foo()
	fmt.Printf("*y = %d\n", *y)
}

func printDemo() {
	name := "Bojan"
	fmt.Printf("Hello," + name + "\n")
}

func conversionDemo() {
	var x int32 = 1
	var y int16 = 2

	x = int32(y)
	fmt.Printf("x = %d", x)

	var a, b int = 3, 4
	avg := float64(a+b) / 2
	fmt.Println("avg = ", avg)
	avg = float64(a+b) / 2.0
	fmt.Println("avg = ", avg)
	avg = float64(float64(a+b) / 2.0)
	fmt.Println("avg = ", avg)
}

func convStringToNumberDemo() {
	i, _ := strconv.Atoi("10")
	y := i * 2
	fmt.Println(y)
}

func constDemo() {
	const x = 1
	const (
		y = 2
		s = "Bojan"
	)
}

func forLoopDemo() {
	for i := 0; i < 10; i++ {
		fmt.Println("i = %d", i)
		if i == 5 {
			break
		}
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Printf("j = %d", i)
	}

	j := 0
	for j < 10 {
		fmt.Printf("j = %d", j)
		j++
	}
	for {
		fmt.Printf("infinite loop")
	}
}

func forLoopDemo2() {
	var xtemp int
	x1 := 0
	x2 := 1
	for x := 0; x < 5; x++ {
		xtemp = x2
		x2 = x2 + x1
		x1 = xtemp
	}
	fmt.Println(x2)
}

func switchDemo() {
	x := 2
	switch x {
	case 1:
		fmt.Printf("case 1")
	case 2:
		fmt.Printf("case 2")
	default:
		fmt.Printf("default")
	}
}

func taglessSwitchDemo() {
	x := 2
	switch {
	case x < -1:
		fmt.Printf("case 1")
	case x > 1:
		fmt.Printf("case 2")
	default:
		fmt.Printf("default")
	}

	fmt.Println("\n2nd example: x == 7:")

	x = 7
	switch {
	case x > 3:
		fmt.Printf("1")
	case x > 5:
		fmt.Printf("2")
	case x == 7:
		fmt.Printf("3")
	default:
		fmt.Printf("4")
	}
}

func inputNumber() {
	var n int
	fmt.Printf("Type in some number: ")
	_, err := fmt.Scan(&n)
	if err == nil {
		fmt.Println("n = ", n)
	}
}

func scanStringDemo() {
	var str string
	fmt.Print("scanStringDemo(): Type in some string: ")
	fmt.Scan(&str)
	fmt.Println("Typed string is: ", str)
}

func scanlnStringDemo() {
	var str string

	fmt.Print("scanlnStringDemo(): Type in some string: ")
	fmt.Scanln(&str) // ENTER terminates input
	fmt.Println("Typed string is: ", str)
}

func stringsDemo() {
	s := strings.Replace("ianianian", "ni", "in", 2) // expected: iainainan
	fmt.Println(s)
}

func scanDemo() {
	//inputNumber()
	//scanStringDemo()
	//scanfStringDemo()
	scanlnStringDemo()
}

// func pointersDemo() {
// 	var x int
// 	var y *int
// 	z := 3
// 	y = &z
// 	//x = &y
// }

func main() {
	// conversionDemo()
	// usePointerToFunctionLocalVariable()
	// printDemo()
	// iotaDemo()
	// scanDemo()
	convStringToNumberDemo()
	// stringsDemo()
	// taglessSwitchDemo()
	forLoopDemo2()
}
