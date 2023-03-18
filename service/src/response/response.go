package response

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
)

type Header struct {
	ID       uint32
	DataSize uint32
}

func SendMessage(outgoing chan []byte, id uint32, jsonData map[string]interface{}) {
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	header := Header{
		ID:       id,
		DataSize: uint32(len(jsonBytes)),
	}

	headerBytes := new(bytes.Buffer)
	err = binary.Write(headerBytes, binary.BigEndian, header)
	if err != nil {
		panic(err)
	}

	packetSize := len(headerBytes.Bytes()) + len(jsonBytes) + 4
	packetBytes := make([]byte, packetSize)
	binary.BigEndian.PutUint32(packetBytes[:4], uint32(packetSize))
	binary.BigEndian.PutUint32(packetBytes[4:8], id)
	copy(packetBytes[8:], jsonBytes)

	outgoing <- packetBytes
}
