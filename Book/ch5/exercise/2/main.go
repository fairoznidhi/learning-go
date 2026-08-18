package main

import (
	"fmt"
	"os"
	"io"
)

func fileLen(filename string)(int, error){
	file, err :=os.Open(filename)
	
	if err != nil {
		return 0, err
	}
	// defer file.Close()
	defer func() {
		fmt.Println("closing file")
		file.Close()
	}()

	byteCount,err := io.ReadAll(file)
	// fmt.Println("File:", byteCount)
	fmt.Println("File use finished")
	return len(byteCount), err
}

func main(){
	data, err := fileLen("file1.txt")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File length:", data)
	}
}