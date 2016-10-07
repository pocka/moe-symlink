package main

import (
	//"os"
	"fmt"
	"path/filepath"
)

func main() {

	files, _ := filepath.Glob("*_*_")

	fmt.Println(files)
}
