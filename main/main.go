package main

import (
	"fmt"
	"github.com/oking02/PhotoImporter/setstaff"
	"github.com/oking02/PhotoImporter/configurations"
	"os"
)

func main() {

	fmt.Println("Hello World!")

	args := os.Args[1:]

//	argslength := len(args)

	configs := configurations.GetConfigs()

	fmt.Println(configs)

	switch args[0] {

	case "get":
		fmt.Println("Get Staff")
		break
	case "set":
		setstaff.SetStaff(args[1:], configs)
		break
	default:
		fmt.Println("Default")

	}

}
