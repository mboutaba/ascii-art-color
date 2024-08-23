package op

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func HandleArgs() ([]string , string, string, string){

	var resultInput []string
    var textInput string
	var pattern string
	var color string
	var banner string

	regChar := regexp.MustCompile(`[ -~]+`)

	regFlag := regexp.MustCompile(`^--color=red$|^--color=green$|^--color=blue$`)

    // regOutpout := regexp.MustCompile(``)

	/////////////////////////////////////////////////////////
	if len(os.Args) < 2 || len(os.Args) > 5 || (!regFlag.MatchString(os.Args[1]) && len(os.Args) > 3) {

		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println()
		fmt.Println("EX: go run . --color=<color> <substring to be colored> \"something\" ")

		os.Exit(1)

	}

    /////////////////////////////////////////////////////////////
		
 if len(os.Args) > 3 && regFlag.MatchString(os.Args[1]) {       // color + pattern + word + ...
  
         colorFlag := os.Args[1]

		 colorElements := strings.Split(colorFlag, "=")

		 color = colorElements[1]

		 
		 
		 textInput = os.Args[3]

	     pattern = os.Args[2]

		 if len(os.Args) == 5 {

			banner = os.Args[4]
		 }
         
		 
   }else if len(os.Args) == 3 && regFlag.MatchString(os.Args[1]) {    // color + "word"

         colorFlag := os.Args[1]

         colorElements := strings.Split(colorFlag, "=")

         color = colorElements[1]


	     textInput = os.Args[2]

	     pattern = ""


   }else if len(os.Args) == 2 && !regFlag.MatchString(os.Args[1]){    // only word       

        textInput = os.Args[1]

		// pattern = ""


   }else if len(os.Args) == 3 && !regFlag.MatchString(os.Args[1]) {   //word + banner

	  textInput = os.Args[1]

      banner = os.Args[2]


   }       
	
   /////////////////////////////////////////////////////////////

		




		wordsInput := strings.Split(textInput, "\\n")

		for i := 0; i < len(wordsInput); i++ {

			if wordsInput[i] == "" {

				resultInput = append(resultInput, "")
				continue

			}

			if !regChar.MatchString(wordsInput[i]) && wordsInput[i] != "" {

				fmt.Println("invalide word")

				os.Exit(1)

			}

			splited := strings.Split(wordsInput[i], "")

			for i, v := range wordsInput[i] {
				if (v < 32 || v > 126) && i < len(splited)-1 {

					fmt.Println("invalid char")
					os.Exit(1)

				}
			}

			resultInput = append(resultInput, strings.Join(splited, ""))

			

		}

		isEmpty := true

		for _, v := range resultInput {

			if v != "" {
				isEmpty = false
			}

		}

		if isEmpty {

			resultInput = resultInput[1:]

		}

	
		return resultInput , color , pattern , banner

}
