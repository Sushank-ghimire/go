package checker

import (
	"fmt"
	"net"
	"time"
)

func Check(domain string, port string) string {
	address := domain + ":" + port
	timeout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp4", address, timeout)

	var status string
	if err != nil {
		status = fmt.Sprintf("Server is down. %v is not reachable.\nError: %v", address, err)
	} else {
		status = fmt.Sprintf("Server is running at %v\nYour Address: %v\nStatus: %v", address, conn.LocalAddr(), conn.RemoteAddr())
		conn.Close() // Close the connection to prevent resource leak
	}
	return status
}
