package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	input := "notes"
	output := "output"
	fmt.Printf("Extracting Boost Note *.json files from directory '%v' into directory '%v'...\n\n", input, output)

	files, err := os.ReadDir(input)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	for _, file := range files {
		filename := fmt.Sprintf("%v/%v", input, file.Name())
		fmt.Println("Reading", filename)
		note := read(filename)
		note.Write(output)
	}
}

type Note struct {
	Title   string
	Content string
}

func read(filename string) Note {
	byteValue, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	var note Note
	err = json.Unmarshal(byteValue, &note)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return note
}

func (note Note) Write(dir string) {
	filePath := fmt.Sprintf("%v/%v.md", dir, note.Title)
	err := ioutil.WriteFile(filePath, []byte(note.Content), 0666)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
