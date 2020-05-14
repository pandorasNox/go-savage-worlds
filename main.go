package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pandorasNox/go-savage-worlds/pkg/savage"
	yaml "gopkg.in/yaml.v2"
)

func main() {
	character, err := characterFromStdin()
	if err != nil {
		log.Fatalf("can't get character: %s", err)
	}

	// fmt.Printf("%+v", character)
	PrettyPrint(character)
}

func PrettyPrint(data interface{}) {
	var p []byte
	//    var err := error
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}

func characterFromStdin() (savage.Character, error) {
	info, err := os.Stdin.Stat()
	if err != nil {
		return savage.Character{}, fmt.Errorf("can't read info from Stdin: %s", err)
	}
	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		errorMsg := "The command is intended to work with pipes.\n"
		errorMsg += "Usage:\n"
		errorMsg += "  cat file | savage"
		return savage.Character{}, fmt.Errorf(errorMsg)
	}

	character, err := LoadCharacter(os.Stdin)
	if err != nil {
		return savage.Character{}, fmt.Errorf("can't read character yaml from Stdin: %s", err)
	}

	return character, nil
}

// LoadCharacter load the char from the reader.
func LoadCharacter(r io.Reader) (savage.Character, error) {
	d := yaml.NewDecoder(r)
	d.SetStrict(true)

	cfg := savage.Character{}

	err := d.Decode(&cfg)
	if err != nil {
		return savage.Character{}, err
	}

	return cfg, nil
}
