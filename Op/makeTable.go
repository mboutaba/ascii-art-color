package op

import (
	"bufio"
	//"fmt"
	"os"
)

func Maketable(Banner string) [][]string {
	
     if Banner == "" {

		Banner = "standard"
	 }

	input, _ := os.Open(Banner + ".txt")
	// scan the banner.txt file for creating a slice of allLines
	scanner := bufio.NewScanner(input)

	var allLines []string

	var charSlices [][]string

	for scanner.Scan() {

		text := scanner.Text() // creation of the slice of lines

		allLines = append(allLines, text)

	}

	y := 0

	for i := 1; i < len(allLines); i++ {

		if len(charSlices) <= y {
			charSlices = append(charSlices, []string{})
		}
		// skiping empty spaces and appending the lines 8 by 8 in the "charSlices" table
		for x := 0; x < 8 && i < len(allLines); x++ {

			charSlices[y] = append(charSlices[y], allLines[i])
			i++
		}

		y++
	}

	return charSlices
}
