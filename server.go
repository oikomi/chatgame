

package main

import (
	"./config"
	"./server"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(0)
	}
	cfg, err := config.LoadConfig(os.Args[1])
	//log.Fatalln(cfg.Listen)
	if err != nil {
		log.Fatalln(err.Error())
		return 
	}
	
	server := server.CreateServer()
	
	server.Listen(cfg.Listen)
	
	
}