package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]

	configs := GetConfigs()

	switch args[0] {

	case "get":
		getStaff(configs)
		break
	case "set":
		setStaff(args[1:], configs)
		break
	case "delete":
		deleteStaff(configs, args)
		break
	default:
		fmt.Println("Default")

	}
}
