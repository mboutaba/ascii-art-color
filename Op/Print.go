package op

import (
	"fmt"
	"strings"
)

func Print(wordsInput []string, charSlices [][]string, color string, pattern string) {
	
	var Reset = "\033[0m"
	var Green = "\033[32m"
	var Red = "\033[31m"
	var Blue = "\033[34m"

	Color := ""

	if color == "red" {
   
		Color = Red  //"\033[31m"

	     }else if color == "green" {

			Color = Green  //"\033[32m"

	        }else if color == "blue" {

				Color = Blue   //"\033[34m"


	}

	for _, word := range wordsInput {

		if word == "" {
			fmt.Println()
			continue
		}



		
		if pattern != "" && pattern != " " && color != "" {

			for i := 0; i < 8; i++ {
				start := 0
				highlighted := ""

				// Find all occurrences of the pattern within the word
				for {
					index := strings.Index(word[start:], pattern)
					if index == -1 {
						highlighted += word[start:]
						break
					}

					before := word[start : start+index]
					match := word[start+index : start+index+len(pattern)]

					highlighted += before + Color + match + Reset
					start = start + index + len(pattern)
				}

				// Print the highlighted text
				inSequence := false

				for _, vv := range highlighted {
                      if vv == '\033' {
						inSequence = true
					}
					if inSequence {
						fmt.Print(string(vv))
						if vv == 'm' {
							inSequence = false
						}

					} else {
					if vv >= 32 && vv <= 126 {
							fmt.Print(charSlices[vv-32][i])

						}
				
					}
				}
				fmt.Println()
			}


		}else if ( pattern == "" ||  pattern == " ") && color != "" {
  
			for i := 0; i < 8; i++ {

				for _, v := range word {

					if v >= 32 && v <= 126 {
						fmt.Print( Green  +charSlices[v-32][i] + Reset)
					}

				}
				fmt.Println()
			}



		} else {

			for i := 0; i < 8; i++ {

				for _, v := range word {

					if v >= 32 && v <= 126 {
						fmt.Print(charSlices[v-32][i])
					}

				}
				fmt.Println()
			}

		}
	//	fmt.Println()
	}

}
