package main

import (
	"fmt"
	"github.com/oking02/PhotoImporter/configurations"
	"github.com/oking02/PhotoImporter/deletestaff"
	"github.com/oking02/PhotoImporter/getstaff"
	"github.com/oking02/PhotoImporter/setstaff"
	"os"
)

func main() {

	args := os.Args[1:]

	configs := configurations.GetConfigs()

	switch args[0] {

	case "get":
		getstaff.GetStaff(configs)
		break
	case "set":
		setstaff.SetStaff(args[1:], configs)
		break
	case "delete":
		deletestaff.DeleteStaff(configs, args)
		break
	default:
		fmt.Println("Default")

	}
}
