package main

import (
	"fmt"
	"os"
	"github.com/oking02/PhotoImporter/setstaff"
)

func main()  {

	fmt.Println("Hello World!")

	args := os.Args[1:]

	argslength := len(args)
	fmt.Println("Number of arguments: ", argslength)

	switch args[0] {

	case "get":
		fmt.Println("Get Staff")
		break
	case "set":
		setstaff.SetStaff(args[1:])
		break
	default:
		fmt.Println("Default")

	}

}


