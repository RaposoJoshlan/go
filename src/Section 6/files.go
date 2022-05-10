package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var newFile *os.File
	fmt.Printf("%T\n", newFile)

	var err error
	newFile, err = os.Create("a.txt")
	if err != nil {
		// fmt.Println(err)
		// os.Exit(1)
		log.Fatal(err)
	}

	err = os.Truncate("a.txt", 0)
	if err != nil {
		log.Fatal(err)
	}

	newFile.Close()

	file, err := os.OpenFile("a.txt", os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		log.Fatal(err)
	}

	file.Close()

	var fileInfo os.FileInfo
	fileInfo, err = os.Stat("a.txt")

	p := fmt.Println
	p("File name:", fileInfo.Name())
	p("Size in bytes:", fileInfo.Size())
	p("Last Modified:", fileInfo.ModTime())
	p("Is Directory:", fileInfo.IsDir())
	p("Permissions:", fileInfo.Mode())

	fileInfo, err = os.Stat("b.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
		}
	}

	// oldPath := "a.txt"
	// newPath := "aaa.txt"

	// err = os.Rename(oldPath, newPath)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = os.Remove("aaa.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//new func
	writingBytesToFile()

}

func writingBytesToFile() {
	file, err := os.OpenFile(
		"b.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	byteSlice := []byte("I learn Golang!")

	byteWritten, err := file.Write(byteSlice)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Byte written: %d\n", byteWritten)

	bs := []byte("Go programming is cool")
	err = ioutil.WriteFile("c.txt", bs, 0644)
	if err != nil {
		log.Fatal(err)
	}

	//New func
	bufferedWritterFunc()

}

func bufferedWritterFunc() {
	file, err := os.OpenFile("myFile.txt", os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	bufferedWriter := bufio.NewWriter(file)

	bs := []byte{97, 98, 99}
	bytesWritten, err := bufferedWriter.Write(bs)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Bytes written to buffer (not file) %d\n", bytesWritten)

	bytesAvailable := bufferedWriter.Available()
	log.Printf("Bytes availbale in buffer %d\n", bytesAvailable)

	bytesWritten, err = bufferedWriter.WriteString(
		"\nJust a random string",
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written: %d\n", bytesWritten)

	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes buffered %d\n", unflushedBufferSize)

	bufferedWriter.Flush()

	bufferedWriter.Reset(bufferedWriter)

	//new func
	readingBytesFromAFile()
}

func readingBytesFromAFile() {
	file, err := os.Open("test.Txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	bs := make([]byte, 2)

	numberBytesRead, err := io.ReadFull(file, bs)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Number of bytes read: %d\n", numberBytesRead)
	log.Printf("Data read: %s\n", bs)

	file, err = os.Open("files.go")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data as string: %s\n", data)
	fmt.Printf("Number of bytes read: %d\n", len(data))

	data, err = ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data read: %s\n", data)

	readingFileLineByLineUsingScanner()
}

func readingFileLineByLineUsingScanner() {
	file, err := os.Open("my_file.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	//scanner.Split(bufio.ScanRunes)

	success := scanner.Scan()
	if success == false {
		err = scanner.Err()
		if err == nil {
			log.Println("Scan completed. Reached end of file")
		} else {
			log.Fatal(err)
		}
	}

	fmt.Println("First line found:", scanner.Bytes())

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal()
	}

	//New func
	scanningForUserInputs()
}

func scanningForUserInputs() {
	scanner := bufio.NewScanner(os.Stdin)

	// fmt.Printf("%T\n", scanner)

	scanner.Scan()

	text := scanner.Text()
	bytes := scanner.Bytes()

	fmt.Println("Input Text:", text)
	fmt.Println("Input Bytes:", bytes)

	for scanner.Scan() {
		text = scanner.Text()
		fmt.Println("You entered:", text)

		if text == "exit" {
			fmt.Println("Exiting the scanning....")
			break
		}
	}

	fmt.Println("Just a message after the for loop")
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
