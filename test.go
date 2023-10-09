package main

import (
	"fmt"
)

func negateInt(x int) int{
	return x*-1
}

func testFunc(x string) func(){
	fmt.Print(x + " Frank")
	return func(){ fmt.Println(" Sinatra")}
}

func testPointer(str *string){
	*str = "changed"
}

type Frank struct{
	name string
	age int
}

func test(){
	var yeet = "yeetxd"
	test := 5
	fmt.Println(yeet + string(test))

	var num1 float64 = 20.5
	var num2 int = 5

	var num3 = num1/float64(num2)
	fmt.Println(num3)

	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan();
	// input := scanner.Text()
	// fmt.Printf("you've been caught %sing\n\n", input)

	fmt.Println(negateInt(5))

	testFunc("Gone Like")()

	x := 8
	y := &x
	fmt.Println("x: %v, y: %v", x, *y)

	text := "some text"
	var pointer *string = &text
	fmt.Printf("%v, %v\n",pointer, *pointer)

	sinatra := Frank{name:"Sinatra"}
	sinatra.age = 5
	fmt.Println(sinatra)
}

