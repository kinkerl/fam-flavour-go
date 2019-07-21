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

type flavouraddonmeta struct {
	Name string `yaml:"name"`
}

type flavouraddon struct {
	Meta flavouraddonmeta `yaml:"meta"`
}

type flavourproject struct {
	Addons yaml.MapSlice `yaml:"addons"`
}

func main() {

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
	new_addon := flavouraddon{}
	err := yaml.Unmarshal([]byte(addon_yml), &new_addon)
	check(err)
	fmt.Printf("--- new addon ---:\n%v\n\n", new_addon)

	// read flavour file
	flavour := flavourproject{}
	flavour_yml, err := ioutil.ReadFile("flavour.yml")
	check(err)
	err2 := yaml.Unmarshal([]byte(flavour_yml), &flavour)
	if err2 != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- flavour ---:\n%v\n\n", flavour)

	new_thing := make(map[string]interface{})
	new_thing["manager"] = os.Getenv("FAM_IDENTIFIER")

	flavour.Addons = append(flavour.Addons, yaml.MapItem{new_addon.Meta.Name, new_thing})

	s, err := yaml.Marshal(&flavour)
	errs := ioutil.WriteFile("flavour_new.yml", s, 0644)
	check(errs)

}
