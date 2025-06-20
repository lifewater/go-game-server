package main

import (
	"fmt"
	"server/pkg/packets"

	"google.golang.org/protobuf/proto"
)

func main() {
	// create chat packet
	packet := &packets.Packet{
		SenderId: 69,
		Msg:      packets.NewChatMessage("Chat Message!"),
	}

	fmt.Println(packet)

	// turn protobuf packet into primary data (aka convert packet to byte slice [byte])
	data, err := proto.Marshal(packet)
	if err != nil {
		fmt.Println("Error marshaling packet: ", err)
		return
	}
	fmt.Println(data)
}
