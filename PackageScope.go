package main

//note file that is connected and being used
//is PackScope2.go

//Note that both files must be run at same time
//to negate errors and run successfully
import "fmt"

//Note go automatically makes connections to different files in same package

func main() {
	//retreiving sayHello function for other file
	sayHello("mario")

	for _, v := range points {
		fmt.Println(v)
	}
}
