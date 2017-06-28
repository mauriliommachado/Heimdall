package main

import (
	"./server"
)

func main() {
	server.Start(server.ServerProperties{Address: "/goid", Port: "8080"})
}
