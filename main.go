package main

import (
	"fmt"
	"os"
)

var (
	file_path = "D:\\Descargas\\enron_mail_20110402\\enron_mail_20110402\\maildir"
)

func read_file(file_path string) {
	fmt.Printf("Leyendo %v\n", file_path)
}

// func list_dir() {
// 	files, err := os.ReadDir(file_path)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	fmt.Println(len(files))
// 	// fmt.Println(files)

// 	for i, file := range files {
// 		fmt.Println(i, file.Name(), file.IsDir())
// 	}
// }

func inspectDirectory(path string, currentFile int, countFiles int, isDir bool) {
	fmt.Println(path, currentFile, countFiles, isDir)
	if isDir == false {
		read_file(path)
		return
	}
	if currentFile == countFiles {
		return
	}
	files, err := os.ReadDir(path)
	if err != nil {
		return
	}
	// fmt.Println(len(files))
	inspectDirectory(path + "\\" + files[currentFile].Name(), currentFile + 1, len(files), files[currentFile].IsDir())
	fmt.Println(path, currentFile, countFiles, isDir)

}

func inspectAllDirectory(path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		return
	}
	// fmt.Println(len(files))
	inspectDirectory(path, 0, len(files), files[0].IsDir())
	// fmt.Println(len(files))
}

func main() {
	// read_file("Hola")
	// list_dir()
	inspectAllDirectory(file_path)
}