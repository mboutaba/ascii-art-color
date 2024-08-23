package main

import (
	//"fmt"
	//"os"

	//"fmt"
	op "main/Op"
)

func main() {

	wordsInput, color, pattern, banner := op.HandleArgs() // handlings args

	charSlices := op.Maketable(banner) // creating the [][] table of ascii

	op.Print(wordsInput, charSlices, color, pattern) // Printing the table of ascii

	// }
}

//////   EX: go run . --color=<color> <substring to be colored> "something"

////    go run . --color=red kit "a king kitten have kit" shadow
