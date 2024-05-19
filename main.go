package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// go run . [STRING] [BANNER]
	InputFile, text := CheckFormat()

	if text == "" {
		
			return
		
	}
	
	slice, slicedArgs := FormatofBanner(InputFile, text)
	DrawAsciiFS(slice, slicedArgs)
}

func CheckFormat() (string, string) {
	slice := os.Args[1:]
	if len(os.Args[1:]) > 2 {
		log.Fatalln("Usage: go run . [STRING] [BANNER] \nEX: go run . \"something\"")
	}
	var InputFile string
	var text string
	for i := range slice {
		if slice[i] == "standard" || slice[i] == "standard.txt" || slice[i] == "shadow" || slice[i] == "shadow.txt" || slice[i] == "thinkertoy" || slice[i] == "thinkertoy.txt" {
			if slice[i] == "standard" || slice[i] == "shadow" || slice[i] == "thinkertoy" {
				InputFile = slice[i] + ".txt"
			} else {
				InputFile = slice[i]
			}
		} else {
			text = slice[i]
		}
	}
	if text == "" {
		if InputFile == "" {
		return "" , ""
		}else{
			log.Fatalln("Usage: go run . [STRING] [BANNER] \nEX: go run . \"something\"")
		}
	}
	if InputFile == "" {
		InputFile="standard.txt"
	}
	return InputFile, text
}

func FormatofBanner(InputFile, text string) ([]string, []string) {
	var sep string
	if InputFile == "standard.txt" || InputFile == "shadow.txt" {
		sep = "\n"
	} else {
		sep = "\r\n"
	}
	// Traite the file
	data, err := os.ReadFile(InputFile)
	if err != nil {
		log.Fatalln(err)
	}
	slice := RemoveEmptyStrings(strings.Split(string(data), sep))
	slicedArgs := strings.Split(text, `\n`)
	return slice, slicedArgs
}

func DrawAsciiFS(slice, slicedArgs []string) {
	var Result string
	for _, word := range slicedArgs {
		if word != "" {
			for i := 0; i < 8; i++ {
				for _, char := range word {
					if char < 32 || char > 126 {
						log.Fatalln("You did entered an inprintabale character !!!")
					} else {
						start := int(char-32)*8 + i
						Result += slice[start]
					}
				}
				Result += "\n"
			}
		} else {
			Result += "\n"
		}
	}
	if IsAllNewLines(Result) {
		Result = Result[1:]
	}
	fmt.Print(Result)
}

func RemoveEmptyStrings(slice []string) []string {
	var temp []string
	for i := range slice {
		if slice[i] != "" {
			temp = append(temp, slice[i])
		}
	}
	return temp
}

func IsAllNewLines(str string) bool {
	for _, char := range str {
		if char != '\n' {
			return false
		}
	}
	return true
}
