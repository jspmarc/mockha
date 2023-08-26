package utils

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"net"
)

func GetRandomTcpPort() uint16 {
	log.Println("Get a random TCP port")
	randomPort := uint16(rand.Int31n(math.MaxInt16))
	isPortOpen := false

	for !isPortOpen {
		c, e := net.Dial("tcp", fmt.Sprintf("localhost:%d", randomPort))

		if e == nil && randomPort >= 1024 {
			isPortOpen = true
			c.Close()
		} else {
			randomPort = uint16(rand.Int31n(math.MaxInt16))
		}
	}

	return randomPort
}
