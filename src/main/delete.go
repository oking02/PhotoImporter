package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const staffUsed = `Cannot delete staff. Currenty assigned to these wards:`

func deleteStaff(configs Config, args []string) {

	client := &http.Client{}

	uniqueString := args[1]

	url := buildDeleteUrl(configs, uniqueString)

	r, _ := http.NewRequest("GET", url, nil)

	r.SetBasicAuth(configs.User, configs.Password)

	resp, err := client.Do(r)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode == 406 {
		fmt.Println("\n")
		fmt.Println("\t", staffUsed)
		fmt.Println("\t\t", string(body))

	}

	if resp.StatusCode == 200 {

		fmt.Println("Successfuly Deleted!")

	}

}

func buildDeleteUrl(configs Config, uniqueString string) string {

	var buffer bytes.Buffer

	buffer.WriteString(configs.BaseUrl)
	buffer.WriteString("delete/")
	buffer.WriteString(uniqueString)

	return buffer.String()

}
