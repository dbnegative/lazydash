package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
)

//LoadFromFile loads prometheus metrics fromfile
func LoadFromFile(path string) ([]byte, error) {
	promFile, err := os.Open(path)
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}

	b, err := ioutil.ReadAll(promFile)

	if err != nil {
		log.Fatalf("could not read from file: %v", err)
	}

	return b, nil
}

//LoadFromStdin loads prometheus metrics from STDIN
func LoadFromStdin() []byte {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		return nil
	}

	reader := bufio.NewReader(os.Stdin)
	var output []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	return ([]byte(string(output)))
}
