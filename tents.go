package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("data/tents.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	type Tent struct {
		Capacity int
		Size [2]int
		Slug, Name, Description, Image, Colors, Content string
	}
	dec := json.NewDecoder(jsonFile)
	_ = os.MkdirAll("content/tents/", 0775)

	// read open bracket
		t, err := dec.Token()
		if err != nil {
			fmt.Println(err)
			fmt.Println(t)
		}

		// while the array contains values
		for dec.More() {
			var a Tent
			var out = ""
			// decode an array value (Article)
			err := dec.Decode(&a)
			if err != nil {
				fmt.Println(err)
			}

			//thisId := fmt.Sprintf("%v", a.Id)

			out += fmt.Sprintf("---\n")
			out += fmt.Sprintf("name: |\n  %v\n", a.Name)
			out += fmt.Sprintf("image: %v\n", a.Image)
			out += fmt.Sprintf("colors: %v\n", a.Colors)
			// out += fmt.Sprintf("sizeX: %d\n", a.Size[0])
			// out += fmt.Sprintf("sizeY: %d\n", a.Size[1])
			out += fmt.Sprintf("size: [%d, %d]\n", a.Size[0], a.Size[1])
			out += fmt.Sprintf("capacity: %d\n", a.Capacity)
			out += fmt.Sprintf("---\n")
			out += fmt.Sprintf("%v\n", a.Content)

			file, err := os.Create("content/tents/" + a.Slug + ".md")
			if err != nil {
				fmt.Println(err)
			}
			file.WriteString(out)
			file.Close()
		}
}