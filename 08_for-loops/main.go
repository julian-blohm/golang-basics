package main

import (
	"fmt"
)

func main() {
	// we have initialization (-> i:=1 optional), condition (i <= 3), and the post statement (i++ optional)
	for i := 1; i <= 3; i++ {
		fmt.Println("Hello World ", i)
	}

	// without initialization and post
	var myInt int = 1
	for myInt < 3 {
		fmt.Println(myInt)
		myInt += 1
	}

	/* Break & Continue
	 * Break: ends loop immediately when it is encounterd
	 * Continue: skips current iteration of loop and continues with next iteration
	 */

	for i := 1; i <= 5; i++ {
		if i == 2 {
			continue // will skip this iteration
		}
		if i == 4 {
			break // break out of loop here
		}
		fmt.Println(i)
	}
}
