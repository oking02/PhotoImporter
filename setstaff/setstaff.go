package setstaff

import (
	"fmt"
	"os"
	"log"
	"encoding/base64"
	_"image"
	_"image/jpeg"
	_ "io/ioutil"
	_"strings"
	"image"
	"bytes"
	"image/jpeg"
	"net/http"
)

func SetStaff(args []string ) {

	client := &http.Client{}

	fmt.Println("Set Staff")

	imagePath := args[0]
//	name := args[1]
//	uniqueId := args[2]

	imgBase64 := getImage(imagePath)

	r, _ := http.NewRequest("POST", "http://localhost:8080/WardMonitor/importTool/set/Bob/Bob[Unique]", bytes.NewBufferString(imgBase64))

	r.SetBasicAuth("ward", "monitor")

	_, err := client.Do(r)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Sent")

}

func getImage(path string) string{

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
