package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Staff struct {
	name         string
	uniqueName   string
	lastModified string
}

func getStaff(configs Config) {

	client := &http.Client{}

	url := buildGetUrl(configs)

	// Build Request
	r, _ := http.NewRequest("GET", url, nil)

	r.SetBasicAuth(configs.User, configs.Password)

	resp, err := client.Do(r)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)

	staff := parseReturnString(string(contents))

	for i := 0; i < len(staff); i++ {

		staffStr := strings.Split(staff[i], ",")

		name := staffStr[1]
		name = removerStartAndEndChart(name)

		uniqueId := staffStr[2]
		uniqueId = removerStartAndEndChart(uniqueId)

		lastMod := staffStr[3]
		t := msToTime(lastMod)

		// Format YYYYMMDDHHMMSS
		lastMod = t.Format("20060102150405")

		fmt.Printf("%s|%s|%s\n", uniqueId, lastMod, name)

	}

	fmt.Println(len(staff))

}

func buildGetUrl(configs Config) string {

	var buffer bytes.Buffer

	buffer.WriteString(configs.BaseUrl)
	buffer.WriteString("get")

	return buffer.String()

}

/*

	Currently Backend is returning a toString Java Array List.
	In this format. With more than one obviously. ID , Name, Unique String, Date Last Modified
	[[1, "John Smith", "John Smith[UniqueId]", 12345689]]

*/

func parseReturnString(str string) []string {

	str = str[1:]
	str = str[:len(str)-1]

	staffStrs := strings.Split(str, "],[")
	lastIndex := len(staffStrs) - 1

	// Remove leftover [  from first and ] from last in the list
	staffStrs[0] = removeStartChar(staffStrs[0])
	staffStrs[lastIndex] = removeEndChars(staffStrs[lastIndex])

	return staffStrs

}

func removeStartChar(str string) string {

	str = str[1:]
	return str
}

func removeEndChars(str string) string {

	length := len(str)
	str = str[:length-1]

	return str
}

func removerStartAndEndChart(str string) string {

	str = removeStartChar(str)

	str = removeEndChars(str)

	return str
}

func msToTime(ms string) time.Time {
	msInt, err := strconv.ParseInt(ms, 10, 64)
	if err != nil {
		return time.Time{}
	}

	return time.Unix(0, msInt*int64(time.Millisecond))
}
