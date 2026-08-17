package main

import "fmt"
import "os"
import "log"
import "io"

func main() {
	fmt.Println(os.Args)
	if(len(os.Args)<2){
		log.Fatal("no file specified")
	}
	f,err:=os.Open(os.Args[1])
	if err!=nil{
		log.Fatal(err)
	}
	defer f.Close()
	data := make([]byte, 2048)
	for{
		count,err := f.Read(data)
		if err!=nil{
			if err!=io.EOF{
				log.Fatal(err)
			}
			break
		}
	}
	
}