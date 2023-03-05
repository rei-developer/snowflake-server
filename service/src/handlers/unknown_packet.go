package handlers

import (
	"fmt"
)

func HandleUnknownPacket(id uint32) {
	fmt.Printf("Unknown packet ID: %d\n", id)
}
