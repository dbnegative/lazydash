package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func LoadTestData() Dashboard {

	testdata := "testdash.json"

	jsonFile, err := os.Open(testdata)
	defer jsonFile.Close()

	if err != nil {
		fmt.Println(err)
	}

	b, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Println(err)
	}

	var dash = Dashboard{}
	json.Unmarshal(b, &dash)

	return dash
}
