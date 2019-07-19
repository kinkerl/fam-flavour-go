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

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {

	var new_addon interface{}
	var flavour interface{}

	reader := bufio.NewReader(os.Stdin)
	var addon_yml []byte

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
	fmt.Printf("--- new addon ---:\n%v\n\n", new_addon)

	// read flavour file
	flavour_yml, err := ioutil.ReadFile("flavour.yml")
	check(err)
	err2 := yaml.Unmarshal([]byte(flavour_yml), &flavour)
	if err2 != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- flavour ---:\n%v\n\n", flavour)

	// -------------------------------------------------------------
	// -------------------------------------------------------------
	// -------------------------------------------------------------

	// hier würde ich gerne folgendes machen, pseudocode:
	/*
			flavour["addons"][new_addon["meta"]["name"]] = {
				"manager": "asd"
			}

			ich will also so einen block hinzufügen zu "addon"

			django-divio:
		    	manager: asd

	*/
	new_addon_interface := new_addon.(map[interface{}]interface{})
	addon_meta := new_addon_interface["meta"]
	addon_meta_interface := addon_meta.(map[interface{}]interface{})
	addon_name := addon_meta_interface["name"]

	new_thing := make(map[string]interface{})
	new_thing["manager"] = os.Getenv("FAM_IDENTIFIER")

	m := flavour.(map[interface{}]interface{})
	// TODO
	n := m["addons"].(map[interface{}]interface{})
	n[addon_name] = new_thing
	// -------------------------------------------------------------
	// -------------------------------------------------------------
	// -------------------------------------------------------------

	s, err := yaml.Marshal(&flavour)
	errs := ioutil.WriteFile("flavour_new.yml", s, 0644)
	check(errs)

	//TODO: write string s changes back to flavour.yml

}
