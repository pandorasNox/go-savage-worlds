package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/go-yaml/yaml"
	"github.com/pandorasNox/go-savage-worlds/pkg/savage"
)

func main() {
	character, err := getCharacterFromStdin()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(character)
}

func getCharacterFromStdin() (savage.Character, error) {
	info, _ := os.Stdin.Stat()
	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		errorMsg := "The command is intended to work with pipes.\n"
		errorMsg += "Usage:\n"
		errorMsg += "cat terraform.tfstate | tftoinv"
		return savage.Character{}, fmt.Errorf(errorMsg)
	}

	var inputBuffer bytes.Buffer
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputBuffer.WriteString(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
		return savage.Character{}, fmt.Errorf("fail to read stdin: %s", err)
	}

	// fmt.Printf("inputBuffer:\n%v\n", inputBuffer.String())

	character := &savage.Character{}
	if err := yaml.Unmarshal(inputBuffer.Bytes(), character); err != nil {
		return savage.Character{}, fmt.Errorf("failed to Unmarshal stdin as an terraform.state json: %s", err)
	}

	return *character, nil
}
