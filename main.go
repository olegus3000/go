package main

import (
	"fmt"
	"os"
)

func main() {
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		fmt.Println("miss args")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := output(path, printFiles, "")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func output(path string, includeFiles bool, prefix string) error {
	dir, err := os.Open(path)
	if err != nil {
		return err
	}
	items, err := dir.Readdir(-1)
	for i, item := range items {
		if !includeFiles && !item.IsDir() {
			continue
		}
		var str string
		if len(items)-1 == i || item.IsDir(){
			str = prefix + "└───"
		} else {
			str = prefix + "├───"
		}
		if item.IsDir() {
			fmt.Print(fmt.Sprintf("%s%s\n", str, item.Name()))
		} else {
			fmt.Print(fmt.Sprintf("%s%s %s\n", str, item.Name(), fmt.Sprintf("(%db)", item.Size())))
		}
		if len(items)-1 == i {
			err = output(path+"/"+item.Name(), includeFiles, prefix+"\t")
		} else {
			err = output(path+"/"+item.Name(), includeFiles, fmt.Sprintf("%s│\t", prefix))
		}
		if err != nil {
			return err
		}
	}
	return nil
}

