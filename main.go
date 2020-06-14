package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pandorasNox/go-savage-worlds/pkg/rulebook"
	"github.com/pandorasNox/go-savage-worlds/pkg/savage"
	yaml "gopkg.in/yaml.v2"
)

func main() {
	sheet, err := characterFromStdin()
	if err != nil {
		log.Fatalf("can't get character: %s", err)
	}

	// fmt.Printf("%+v", character)
	// PrettyPrint(sheet)

	rb := rulebook.New(rulebook.SWADE_Attributes, rulebook.SWADE_Skills)

	err = savage.Validate(sheet, rb)
	if err != nil {
		log.Fatalf("sheet is not valid: %s", err)
	}

	fmt.Println("sheet is valid")
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

func characterFromStdin() (savage.Sheet, error) {
	info, err := os.Stdin.Stat()
	if err != nil {
		return savage.Sheet{}, fmt.Errorf("can't read info from Stdin: %s", err)
	}
	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		errorMsg := "The command is intended to work with pipes.\n"
		errorMsg += "Usage:\n"
		errorMsg += "  cat file | savage"
		return savage.Sheet{}, fmt.Errorf(errorMsg)
	}

	character, err := LoadCharacter(os.Stdin)
	if err != nil {
		return savage.Sheet{}, fmt.Errorf("can't read character yaml from Stdin: %s", err)
	}

	return character, nil
}

// LoadCharacter load the char from the reader.
func LoadCharacter(r io.Reader) (savage.Sheet, error) {
	d := yaml.NewDecoder(r)
	d.SetStrict(true)

	cfg := savage.Sheet{}

	err := d.Decode(&cfg)
	if err != nil {
		return savage.Sheet{}, err
	}

	return cfg, nil
}
