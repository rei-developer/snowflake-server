package handlers

import (
	"net"
	"sync"

	"github.com/snowflake-server/src/user"
)

func HandleUserDisconnect(conn net.Conn, users map[uint32]*user.User, mu *sync.Mutex) {
	defer mu.Unlock()
	mu.Lock()
	for id, u := range users {
		print(u.Conn)
		if u.Conn == conn {
			println("유저 종료 : %s %s", u.Email, u.Index)
			delete(users, id)
			break
		}
	}
}
