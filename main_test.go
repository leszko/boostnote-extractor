package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestRead(t *testing.T) {
	inputFile := "notes/note.json"
	note := read(inputFile)

	if note.Title != "Golang resources" {
		t.Errorf("Expected title 'Golang resources', but got '%v", note.Title)
	}

	if note.Content != "Some Content" {
		t.Errorf("Expected content 'Some Content', but got '%v'", note.Content)
	}
}

func TestWrite(t *testing.T) {
	note := Note{Title: "Some Note", Content: "Some Content"}

	dir, _ := ioutil.TempDir("", "")
	note.Write(dir)

	content, _ := ioutil.ReadFile(fmt.Sprintf("%v/Some Note.md", dir))
	if string(content) != "Some Content" {
		t.Errorf("Expected 'Some Content', but got '%v'", content)
	}
}
