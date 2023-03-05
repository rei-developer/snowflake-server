package response

import "encoding/binary"

func SendResponse(outgoing chan []byte, id uint32, message string) {
	response := []byte(message)
	outgoing <- append([]byte{0, 0, 0, byte(id)}, uint32ToBytes(uint32(len(response)))...)
	outgoing <- response
}

func SendSuccessResponse(outgoing chan []byte, message string) {
	SendResponse(outgoing, 1, message)
}

func SendErrorResponse(outgoing chan []byte, message string) {
	SendResponse(outgoing, 2, message)
}

func uint32ToBytes(n uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, n)
	return b
}
