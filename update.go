package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var base = `GoLeet
======

My collection of leet code problems solved in GoLang!

Link to profile, [bhargavsnv100](https://leetcode.com/bhargavsnv100/)

| [Easy](Easy/README.md) | [Medium](Medium/README.md) | [Hard](Hard/README.md) |
|------------------------|----------------------------|------------------------|
| %v                     | %v                         | %v                     |

Problems finished
=================
`

func main() {

	dirs := []string{"Easy", "Medium", "Hard"}
	count := []int{0, 0, 0}

	readmeData := ""

	for i, dir := range dirs {
		readmeData += fmt.Sprintf("\n%s\n%s\n\n", dir, strings.Repeat("-", len(dir)))

		// Remove Problem Directory Readme
		readme := fmt.Sprintf("./%s/README.md", dir)
		os.Remove("./" + dir + "/README.md")

		// Get files in directory
		files, err := ioutil.ReadDir("./" + dir + "/")
		if err != nil {
			log.Fatal(err)
		}

		// Get count of files
		count[i] = len(files)

		// Populate list of solved problems
		data := fmt.Sprintf("%s\n%s\n\n", dir, strings.Repeat("=", len(dir)))
		for _, file := range files {
			row := fmt.Sprintf("- [%s](%s/%s)\n", file.Name(), dir, file.Name())
			data += strings.ReplaceAll(row, dir+"/", "")
			readmeData += row
		}

		// Write problem directory readme
		err = ioutil.WriteFile(readme, []byte(data), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}

	base = fmt.Sprintf(base, count[0], count[1], count[2])
	readmeData = base + readmeData

	ioutil.WriteFile("README.md", []byte(readmeData), 0644)

}
