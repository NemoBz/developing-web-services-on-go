package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println(".")
	dirtree(".", "")
}
func dirtree(path string, prefix string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return
	}

	for i, entry := range entries {

		if i == len(entries)-1 {
			fmt.Printf("%s└── %s\n", prefix, entry.Name())
		} else {
			fmt.Printf("%s├── %s\n", prefix, entry.Name())
		}

		if entry.IsDir() {
			newPath := path + "/" + entry.Name()

			var newPrefix string
			if i == len(entries)-1 {
				newPrefix = prefix + "    "
			} else {
				newPrefix = prefix + "│   "
			}

			dirtree(newPath, newPrefix)
		}
	}
}
