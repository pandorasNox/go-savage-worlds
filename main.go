package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/pandorasNox/go-savage-worlds/pkg/savage"
	yaml "gopkg.in/yaml.v2"
)

func main() {
	character, err := getCharacterFromStdin()
	if err != nil {
		log.Fatalf("can't get character: %s", err)
	}

	fmt.Println(character)
}

func getCharacterFromStdin() (savage.Character, error) {
	info, err := os.Stdin.Stat()
	if err != nil {
		log.Fatalf("can't read info from Stdin: %s", err)
	}
	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		errorMsg := "The command is intended to work with pipes.\n"
		errorMsg += "Usage:\n"
		errorMsg += "  cat file | savage"
		return savage.Character{}, fmt.Errorf(errorMsg)
	}

	yamlFile, err := ioutil.ReadAll(os.Stdin) //ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("can't read \"all\" from Stdin: %s", err)
	}

	character := &savage.Character{}

	err = yaml.Unmarshal(yamlFile, character)
	if err != nil {
		log.Fatalf("can't unmarshal expected yaml input: %s", err)
	}

	return *character, nil
}
