package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir, err := os.Getwd()

	if err != nil {
		fmt.Println(err.Error())
	}

	// fmt.Println(dir)
	//create file and see the file info
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

	//how to create a folder
	// errs := os.Mkdir("kcl", 0777)
	// if errs != nil {
	// 	fmt.Println(errs.Error())
	// }

	base := filepath.Base(dir)
	absulatePath, _ := filepath.Abs("kcl")
	fmt.Println(base, "\n", absulatePath)

	newPath := filepath.Join(absulatePath, "..", "..", "lecture_25")
	fmt.Println(newPath)
	os.Mkdir(newPath, 777)

}

func createFile(fileName, content string) bool {
	posf, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer posf.Close()

	_, err = posf.Write([]byte(content))
	//fmt.Println(n, err)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}
