package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		fmt.Print(len(os.Args))
		//panic("usage go run main.go . [-f]")
	}
	defer fmt.Println(len(os.Args))
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	fmt.Println(path)
	fmt.Println(printFiles)
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("Указанная папка не существует.")
		return err
	}
	//C:\Users\NSmagulov\go\stepik\1\99_hw\tree
	fmt.Println("dssdfs " + path)
	listDir(out, path, printFiles, "")
	return nil
}

func listDir(out io.Writer, path string, printFiles bool, prefix string) error {

	dir, err := os.Open(path)

	if err != nil {
		return err
	}
	defer dir.Close()

	files, err := dir.ReadDir(-1)

	if err != nil {
		return err
	}

	for i, file := range files {
		isLast := i == len(files)-1
		if isLast {
			fmt.Fprintf(out, "%s└───%s\n", prefix, file.Name())
		} else {
			fmt.Fprintf(out, "%s├───%s\n", prefix, file.Name())
		}

		if file.IsDir() {
			newPrefix := ""
			if isLast {
				newPrefix += "    "
			} else {
				newPrefix += "│   "
			}
			err = listDir(out, path+"/"+file.Name(), printFiles, newPrefix)
			if err != nil {
				return err
			}
		}
	}
	return err
}
