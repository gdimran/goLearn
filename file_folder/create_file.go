package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// var (
// 	newFile *os.File
// 	err     error
// )

func main() {
	// newFile, err := os.Create("test2.txt")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// log.Println(newFile)
	// newFile.Close()

	dir, err := os.Getwd()

	if err != nil {
		fmt.Println(err.Error())
	}
	// fmt.Println(dir)
	// //create file and see the file info
	// isErr := createFile("mcwl.txt", "This enhanced version from file content text")
	// fmt.Println(isErr)

	// fi, err := os.Stat("mcwl.txt")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(fi.IsDir())
	// fmt.Println(fi.ModTime().Date())
	// fmt.Println(fi.Name())
	// fmt.Println(fi.Size())

	//create directory folder with permission
	// err := os.Mkdir("mcwl", 0644)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	base := filepath.Base(dir)
	relativePath := filepath.Join("kcl")
	absulatePath, _ := filepath.Abs("mcwl") //how to find absulate path for a folder
	fmt.Println(base)
	fmt.Println(relativePath)
	fmt.Println(absulatePath)

}
