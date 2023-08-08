//create global var and local var and see which var will it take

package main

import "fmt"

var temp1 int = 10 // global scope

func main() {
	temp1 = 20

	fmt.Println(temp1)

	var ex1 string = "Hi" //gobal scope scope

	{
		ex1 = "Hello" //local scope

		fmt.Println(ex1) //so it usually takes the local scope
	}
}
