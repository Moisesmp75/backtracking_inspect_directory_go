package main

import (
	"fmt"
	"os"
	"time"
)

var (
	file_path = "D:\\Descargas\\enron_mail_20110402\\enron_mail_20110402\\maildir"
)

func read_file(file_path string) {
	fmt.Printf("Leyendo %v\n", file_path)
}

func inspectDirectory(path string, isDir bool) {
	fmt.Println(path)
	if !isDir {
		read_file(path)
		return
	}
	files, err := os.ReadDir(path)
	if err != nil {
		return
	}
	for _, file := range files {
		inspectDirectory(path+"\\"+file.Name(), file.IsDir())
	}
}

func main() {
	startTime := time.Now()
	inspectDirectory(file_path, true)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Println("Tiempo transcurrido:", duration)
}
