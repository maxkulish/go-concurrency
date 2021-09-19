package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func connectToService() interface{} {
	time.Sleep(1*time.Second)
	return struct {}{}
}

// startNetworkDaemon creates a server which listen to localhost:8080
func startNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {

		connPool := warmServiceConnCache()

		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("cannot listen: %v", err)
		}
		defer server.Close()

		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("cannot accept connection: %v", err)
				continue
			}
			svcConn := connPool.Get()
			fmt.Fprintln(conn, "")
			connPool.Put(svcConn)
			conn.Close()
		}
	}()
	return &wg
}

// warmServiceConnCache creates a pool of connections
func warmServiceConnCache() *sync.Pool {
	p := &sync.Pool {
		New: connectToService,
	}
	for i := 0; i < 10; i++ {
		p.Put(p.New())
	}
	return p
}

