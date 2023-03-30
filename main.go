package main

import (
	"flag"
	"log"
	"refes/api"
)

func main() {
	address := flag.String("address", "0.0.0.0", "address for the server to listen on")
	port := flag.Int("port", 8080, "port for the server to listen on")

	err := api.Init(address, port)
	if err != nil {
		log.Fatalln(err)
	}
}
