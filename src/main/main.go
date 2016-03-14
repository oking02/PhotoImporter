package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]

	configs := getConfigs()

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
	case "test":
		compress()
		break
	case "dir" :
		resizeAllPhotos("C:\\Users\\ollyking\\Documents\\GoTests\\src\\github.com\\oking02\\PhotoImporter\\testdata", "C:\\Users\\ollyking\\Documents\\GoTests\\src\\github.com\\oking02\\PhotoImporter\\testdatadest\\")
	default:
		fmt.Println("Default")

	}
}


