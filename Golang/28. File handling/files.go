package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("example.txt")
	if err != nil {
		// log the error
		panic(err)
	}
	fileInfo, err := f.Stat()
	if err != nil {
		// log the error
		panic(err)
	}
	fmt.Println("File Name:", fileInfo.Name())
	fmt.Println("files or folder", fileInfo.IsDir())
	fmt.Println("File Size:", fileInfo.Size())
	fmt.Println("File Mode:", fileInfo.Mode())
	fmt.Println("Last Modified:", fileInfo.ModTime())
	fmt.Println("Is Readable:", fileInfo.Mode().IsRegular())
	fmt.Println("Is Directory:", fileInfo.IsDir())
	fmt.Println("Is Hidden:", fileInfo.Name()[0] == '.')
	fmt.Println("Is Executable:", fileInfo.Mode()&0111 != 0)
	fmt.Println("Is Symbolic Link:", fileInfo.Mode()&os.ModeSymlink != 0)
	defer f.Close()

	// read file
	f, err = os.Open("example.txt")
	if err != nil {
		// log the error
		panic(err)
	}

	defer f.Close()

	buf := make([]byte, 1024)
	d, err := f.Read(buf)
	if err != nil {
		// log the error
		panic(err)
	}
	fmt.Println("data", d, buf)

}
