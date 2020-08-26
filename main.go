package main

import (
	"fmt"
	"gurban/src/core"
	"os"
)

func main() {
	printInfo()
	args := os.Args[1:]
	//fmt.Println("[*] Args: ", args)
	if len(args) < 1 {
		fmt.Println("[*] Please provide an argument!")
	} else {
		var term string;
		for i := 0; i < len(args); i++ {
			if i == 0 {
				term = args[i];
			} else {
				term = term + "+" + args[i]
			}
		}

		core.GetEntryForTerm(term)
	}
}

func printInfo() {
	fmt.Println("[*] gurban v0.1 📔")
	fmt.Println("[*] quick CLI urban dictionary lookup")
	fmt.Println("[*] GitHub: https://github.com/woodfairy")
	fmt.Println()
}
