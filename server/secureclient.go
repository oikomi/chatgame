
package server

import (
	"log"
	"net"
	"bufio"
	//"fmt"
)

type SecureClient struct {
	conn     net.Conn
	h        *HeartBeat
	reader   *bufio.Reader
	writer   *bufio.Writer
	incoming chan string
	outgoing chan string
	name     string
	Alive    bool
	mutex         sync.Mutex
}


func CreateClient(conn net.Conn) *SecureClient {
	r := bufio.NewReader(conn)
	w := bufio.NewWriter(conn)
	client := &SecureClient {
		conn  : conn,
		h     : NewHeartBeat(60, 60, 100),
		reader: r,
		writer: w,
		incoming : make(chan string),
		outgoing : make(chan string),
	}
	
	return client
}


func (self *SecureClient) ClientEvent() {
	go self.Read()
	go self.Write()
}

func (self *SecureClient) Read() {
	//fmt.Println("Read")
	for {
		buf := make([]byte, 10000)
		_ , err := self.reader.Read(buf)
		if err != nil {
			return
		}
		self.incoming <- string(buf)
		self.setAlive(true)
	}

}

func (self *SecureClient) Write() {
	//fmt.Println("begin Write")
	for data := range self.outgoing {
		//fmt.Println("Write")
		//_, err := self.writer.WriteString(data + "\n")
		_, err := self.writer.Write([]byte(data))
		
		if err != nil {
			log.Printf("Write error: %s\n", err)
			return
		}
		
		if err := self.writer.Flush(); err != nil {
			log.Printf("Write error: %s\n", err)
			return
		}
	}
}

func (self *SecureClient) setAlive(flag) {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	self.Alive = flag
}

func (self *SecureClient) Close() {
	self.conn.Close()
}

func (self *SecureClient) GetName() string {
	return self.name
}

func (self *SecureClient) SetName(name string) {
	self.name = name
}
