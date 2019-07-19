package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// cat ../flavour/libflavour/libflavour/test/data/example_addon_data_1.yaml | go run add.go

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type T struct {
	META struct {
		NAME    string
		VERSION string
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {

	new_addon := map[string]interface{}{}
	flavour := map[string]interface{}{}

	reader := bufio.NewReader(os.Stdin)
	var addon_yml []byte
	//var flavour_yml []byte

	//read from stdin
	for {
		input, err := reader.ReadByte()
		if err != nil && err == io.EOF {
			break
		}
		addon_yml = append(addon_yml, input)
	}

	// read input go file for addon
	err := yaml.Unmarshal([]byte(addon_yml), &new_addon)
	check(err)
	fmt.Printf("--- t:\n%v\n\n", new_addon)

	// read flavour file
	flavour_yml, err := ioutil.ReadFile("flavour.yml")
	check(err)
	err2 := yaml.Unmarshal([]byte(flavour_yml), &flavour)
	if err2 != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", flavour)

	// -------------------------------------------------------------
	// -------------------------------------------------------------
	// -------------------------------------------------------------

	flavour["addons"]["asd"] = "asd"

	// -------------------------------------------------------------
	// -------------------------------------------------------------
	// -------------------------------------------------------------

	s, err := yaml.Marshal(&flavour)
	fmt.Println(string(s))
	fmt.Println("PATH:", os.Getenv("PATH"))
	//write changes to flavour.yml

}
