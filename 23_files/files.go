package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// fileInfo, err := f.Stat()

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("File name",fileInfo.Name())
	// fmt.Println("File or folder",fileInfo.IsDir())
	// fmt.Println("Size",fileInfo.Size())
	// fmt.Println("File modified at",fileInfo.ModTime())

	// Read file

	// buf := make([]byte, fileInfo.Size())

	// d, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }

	// for i :=0; i < len(buf); i++{
	// fmt.Println("data", d, string(buf[i]))
	// }

	// Whole file read
	// f, err := os.ReadFile("example.txt")

	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(f))

	// Read Directory
	// dir, err := os.Open(".")
	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(-1)
	// 	if err != nil {
	// 	panic(err)
	// }
	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name())
	// }

	// create a file
	// f, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// f.WriteString("print hii go")
	// f.WriteString("update")

	// using byte
	// bytes := []byte("Hello golang")
	// n, err := f.Write(bytes)

	// fmt.Println(n)

	// read and write to another file (streaming fashion)

	sourceFile, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	defer sourceFile.Close()

	destFile, err := os.Create("example2.txt")
	if err != nil {
		panic(err)
	}

	defer destFile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}
		e := writer.WriteByte(b)
		if e != nil {
			panic(e)
		}
	}

	writer.Flush()
	fmt.Println("Written to new file successfuly")
}
