package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"os"
)

func setStaff(args []string, configs Config) {

	client := &http.Client{}

	imagePath := args[0]
	name := args[1]
	uniqueId := args[2]

	//Build url from args and config base url
	url := buildSetUrl(configs, name, uniqueId)

	//Get image base 64 from image path argument
	imgBase64 := getImage(imagePath)

	// Build Request
	r, _ := http.NewRequest("POST", url, bytes.NewBufferString(imgBase64))

	//Set Basic Auth using config
	r.SetBasicAuth(configs.User, configs.Password)

	//Send Request
	_, err := client.Do(r)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sent")

}

func buildSetUrl(configs Config, name, uniqueId string) string {

	var buffer bytes.Buffer

	buffer.WriteString(configs.BaseUrl)
	buffer.WriteString("set/")
	buffer.WriteString(name)
	buffer.WriteString("/")
	buffer.WriteString(uniqueId)

	return buffer.String()

}

func getImage(path string) string {

	reader, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer reader.Close()

	buffer := new(bytes.Buffer)

	m, _, err := image.Decode(reader)

	if err := jpeg.Encode(buffer, m, nil); err != nil {
		log.Fatal(err)
	}

	return base64.StdEncoding.EncodeToString(buffer.Bytes())

}
