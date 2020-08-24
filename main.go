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
		core.GetEntryForTerm(args[0])
	}
}

func printInfo() {
	fmt.Println("[*] gurban v0.1 📔")
	fmt.Println("[*] quick CLI urban dictionary lookup")
	fmt.Println("[*] GitHub: https://github.com/woodfairy")
	fmt.Println("[*] Copyright (c) 2020 by woodfairy")
	fmt.Println()
}
