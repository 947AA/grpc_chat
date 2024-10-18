package main

import (
	"fmt"

	cm "github.com/947AA/grpc_chat/pb"
	"google.golang.org/protobuf/proto"
)

func main() {

	fmt.Println("Hello world!")
	Com1 := cm.Cpu{}
	data, _ := proto.Marshal(&Com1)
	var sync []*cm.Cpu
	com2 := cm.Cpu{}
	proto.Unmarshal(data, &com2)

	sync = append(sync, &com2)
	fmt.Println(sync[0].String())
}
