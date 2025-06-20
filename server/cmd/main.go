package main

import (
	"fmt"
	"server/pkg/packets"

	"google.golang.org/protobuf/proto"
)

func main() {
	// create chat packet
	packet1 := &packets.Packet{
		SenderId: 69,
		Msg:      packets.NewChatMessage("Chat Message!"),
	}
	fmt.Printf("Raw Data Packet:    %s\n", packet1)

	// turn protobuf packet into primary data (aka convert packet to byte slice [byte])
	pdata, err := proto.Marshal(packet1)
	if err != nil {
		fmt.Println("Error marshaling packet: ", err)
		return
	}
	fmt.Printf("Marshal'd Packet:   %v\n", pdata)

	// turn binary data back into protobuf packet
	rdata := []byte{8, 69, 18, 15, 10, 13, 72, 101, 108, 108, 111, 44, 32, 119, 111, 114, 108, 100, 33}
	packet2 := &packets.Packet{}
	err2 := proto.Unmarshal(rdata, packet2)
	if err2 != nil {
		fmt.Println("Error unmarshaling packet:   ", err)
		return
	}
	fmt.Printf("Unmarshal'd Packet: %v\n", packet2)
}
