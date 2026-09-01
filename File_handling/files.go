package main

import (
	"fmt"
	"os"
)

func main() {
	/*f, err := os.Open("example.txt")
	if err != nil {
		//log the error
		//panic
		panic(err)
	}

	fileinfo, err := f.Stat()

	if err != nil {
		panic(err)
	}

	fmt.Println("File name:", fileinfo.Name())
	fmt.Println("File size:", fileinfo.Size())*/

	//read file content
	/*f, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := make([]byte, 12) // dynamic way : fileinfo.Size()

	d, err := f.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bytes read:", d)
	fmt.Println("Content:", string(buf))*/

	// if the file is small, we can use os.ReadFile() to read the entire file content at once
	/*f, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(f))*/

	/*dir, err := os.Open("../")
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	fileinfo, err := dir.ReadDir(-1)
	if err != nil {
		panic(err)
	}

	for _, fi := range fileinfo {
		fmt.Println(fi.Name(), fi.IsDir())
	}*/

	//create a new file
	/*f, err := os.Create("newfile.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()*/

	//write to the new file
	//_, err = f.WriteString("Hello, World!")
	/*f.WriteString("\nHii Go")
	f.WriteString("\nWelcome to Go programming")
	if err != nil {
		panic(err)
	}*/

	/*bytes := []byte("Hello, World!")

	f.Write(bytes)*/

	//read and write to another file (streaming fashion)

	/*sourcefile, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer sourcefile.Close()

	destfile, err := os.Create("example2.txt")
	if err != nil {
		panic(err)
	}
	defer destfile.Close()

	reader := bufio.NewReader(sourcefile)
	writer := bufio.NewWriter(destfile)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}
		err1 := writer.WriteByte(b)
		if err1 != nil {
			panic(err1)
		}
	}
	writer.Flush()

	fmt.Println("File copied successfully")*/

	//delete file

	err := os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("File deleted successfully")
}
