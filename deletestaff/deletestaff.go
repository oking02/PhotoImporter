package deletestaff

import (
	"bytes"
	"fmt"
	"github.com/oking02/PhotoImporter/configurations"
	"io/ioutil"
	"log"
	"net/http"
)

const staffUsed = `Cannot delete staff. Currenty assigned to these wards:`

func DeleteStaff(configs configurations.Config, args []string) {

	client := &http.Client{}

	uniqueString := args[1]

	url := buildUrl(configs, uniqueString)

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

func buildUrl(configs configurations.Config, uniqueString string) string {

	var buffer bytes.Buffer

	buffer.WriteString(configs.BaseUrl)
	buffer.WriteString("delete/")
	buffer.WriteString(uniqueString)

	return buffer.String()

}
