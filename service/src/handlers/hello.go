package handlers

import (
	"github.com/snowflake-server/src/response"
)

func HandleHello(outgoing chan []byte) {
	response.SendSuccessResponse(outgoing, "Hello, client!")

	// Send a response packet to the client
	// resp := Message{Type: loginVerificationResponse, Payload: []byte(`{"success": true}`)}
	// outgoing <- resp.Encode()
}
