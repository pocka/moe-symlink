package main

import (
	"os"
	"fmt"
	"path/filepath"
)

var (
	version string
)

func main() {

	targetFiles := []string{
		"colorcfg.ini",
		"friend.txt",
		"ignore.txt",
		"keyconfig.txt",
		"msgopt.ini",
		"wnd.ini",
		"wnd_docking.ini",
	}

	if _, err := os.Stat("source"); os.IsNotExist(err) {

		fmt.Println("No source directory found.")

		return
	} else {

		for _, file := range targetFiles {

			path := "source/" + file

			if _, err := os.Stat(path); os.IsNotExist(err) {

				fmt.Printf("File not found [%s]\n", path)

				return
			}
		}
	}

	fmt.Printf("moe-symlink %s\n\n", version)

	dirs, _ := filepath.Glob("*_*_")


	if len(dirs) <= 0 {

		fmt.Println("No character data.")

		return
	}

	for _, dir := range dirs {

		fmt.Printf("* Found [%s]\n", dir)

		var proceedsCount int = 0

		for _, file := range targetFiles {

			path := dir + "/" + file

			_, err := os.Readlink(path)

			if err != nil {

				if _, err := os.Stat(path); err == nil {
					fmt.Printf("Delete file [%s]\n", file)

					err := os.Remove(path)

					if err != nil {

						fmt.Printf("Failed to delete [%s]\n", file)

						continue
					}
				}

				symlinkPath, _ := filepath.Abs("source/" + file)

				err := os.Symlink(symlinkPath, path)

				if err != nil {
					fmt.Println("Failed to create link")

					continue
				}

				fmt.Printf("[%s] -> [%s]\n", path, file)

				proceedsCount += 1
			}
		}

		fmt.Printf("Created %d links\n\n", proceedsCount)
	}

	return
}
