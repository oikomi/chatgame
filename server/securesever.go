
package server

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
	"strings"
)

type Message     chan string
type ClientTable map[net.Conn]*SecureClient

type SecureServer struct {
	listener      net.Listener
	clients       ClientTable
	clientcoming  chan net.Conn
	incoming      Message
	
	mutex         sync.Mutex
	scanSessionMutex  sync.Mutex
}

func CreateServer() *SecureServer {
	server := &SecureServer {
		clients      : make(ClientTable),
		clientcoming : make(chan net.Conn),
		incoming     : make(Message),
	}
	
	return server
}

func (self *SecureServer)serverEvent() {
	for {
		select {
		case message := <-self.incoming:
			self.messageProcess(message)
		case conn := <- self.clientcoming:
			self.processClient(conn)
		}
	}
}

func (self *SecureServer)Listen(port string) {
	//log.Printf("connport %s \n", port)

	listener, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalln(err.Error())
        return
    }
	go self.serverEvent()
	go self.scanDeadSession()
	for {
		conn, err := listener.Accept()
		//log.Printf("Accept")
		if err != nil {
			log.Fatalln(err.Error())
			return
		}
		self.clientcoming <- conn
	}
}

func (self *SecureServer)sendToEveryClient(message string) {
	for _, client := range self.clients {
		if client.GetName() == "" {
			client.outgoing <- "you are not logined, Please input your username : \n"
		} else {
			client.outgoing <- message
		}
	}
}

func (self *SecureServer)sendToEveryClientRaw(message string) {
	for _, client := range self.clients {
		if client.GetName() == "" {
			client.outgoing <- "you are not logined, Please input your username : \n"
		} else {
			client.outgoing <- message
		}
	}
}

func (self *SecureServer)sendToSingleClient(client *SecureClient, message string) {
	client.outgoing <- message
}

func (self *SecureServer)processWelcome(client *SecureClient, message string) {
	client.outgoing <- message
}

func (self *SecureServer)messageProcessRaw(message string) {
	self.sendToEveryClientRaw(message)
}

func (self *SecureServer)messageProcess(message string) {
	self.sendToEveryClient(message)
}

func (self *SecureServer)processClient(conn net.Conn) {
	//log.Printf("a new client " + conn.RemoteAddr().String())
	client := CreateClient(conn)
	
	go self.processWelcome(client, "Please input your username : \n")
	
	self.clients[conn] = client
	
	client.ClientEvent()
	go func() {
		for {
			msg := <-client.incoming
			if strings.HasPrefix(msg, "/") {
				if strings.HasPrefix(msg, "/quit") {
					self.delOffline(client)
				} else {
					cmd := CreateCmd()
					msglist := strings.Split(msg, " ")
					cmd.parseCmd(msglist)
					cmd.executeCommand(self, client)
				}
				
			} else {
				//fmt.Println(fmt.Sprintf("%s says: %s", client.GetName(), msg))
				
				self.incoming <- fmt.Sprintf("%s says: %s", client.GetName(), msg)
			}
		}
	}()
}

func (self *SecureServer)getOnline() int {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	
	return len(self.clients)
}

func (self *SecureServer)delOffline(client *SecureClient) {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	self.sendToEveryClientRaw(fmt.Sprintf("%s is offline", client.GetName()))
	client.conn.Close()
	delete(self.clients, client.conn)
}

func (self *SecureServer)scanDeadSession() {
	timer := time.NewTicker(60 * time.Second)
	ttl := time.After(100 * time.Second)
	for {
		select {
		case <-timer.C:
			go func() {
				for _, c := range self.clients {
					self.scanSessionMutex.Lock()
					if c.Alive == false {
						self.delOffline(c)
					} else {
						c.Alive = false
					}
					self.scanSessionMutex.Unlock()
				}
			}()
		case <-ttl:
			break
		}
	}
}