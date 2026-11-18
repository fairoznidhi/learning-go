package main

//go:generate protoc -I=. --go_out=. --go_opt=module=example.com/go-generate-demo --go_opt=Mperson.proto=example.com/go-generate-demo/data person.proto

import (
	"fmt"

	"example.com/go-generate-demo/data"
	"google.golang.org/protobuf/proto"
)

func main() {
	p := &data.Person{
		Name:  "Bob Bobson",
		Id:    20,
		Email: "bob@bobson.com",
	}

	fmt.Println(p)

	protoBytes, err := proto.Marshal(p)
	if err != nil {
		panic(err)
	}

	fmt.Println(protoBytes)

	p2 := &data.Person{}
	err = proto.Unmarshal(protoBytes, p2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p2)
}
