package main

import (
	"os"
	"strings"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

/*
go build
go run
go install
go get

package main vs other

first one is used to create executable files from source code, others are helper codes.
each package main needs a main function to run
*/

func main() {

	// Inefficient approach (strings are immutable in GO, every assignment to s results in copying and creating new string)
	s, sep := "", ""
	for _, arg := range os.Args[:] {
		s += sep + arg
		sep = " "
	}

	// efficient approach
	strings.Join(os.Args[1:], " ")
}
