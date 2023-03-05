package handlers

import (
	"github.com/snowflake-server/src/response"
)

func HandleHello(outgoing chan []byte) {
	response.SendSuccessResponse(outgoing, "Hello, client!")
}
