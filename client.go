package main

import (
    "log"
    "crypto/tls"
	"bufio"
	"os"
	"fmt"
)

func main() {
    log.SetFlags(log.Lshortfile)

    conf := &tls.Config{
        InsecureSkipVerify: true,
    }

    conn, err := tls.Dial("tcp", "127.0.0.1:443", conf)
    if err != nil {
        log.Println(err)
        return
    }
    defer conn.Close()
/*
    n, err := conn.Write([]byte("hello\r\n"))
    if err != nil {
        log.Println(n, err)
        return
    }

    buf := make([]byte, 100)
    n, err = conn.Read(buf)
    if err != nil {
		log.Println("--------1---------")
        log.Println(n, err)
        return
    }

    println(string(buf[:n]))
	
*/

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	//client := CreateClient(conn)

	go func() {
		for {
		    buf := make([]byte, 1000)
		    n, err := conn.Read(buf)
		    if err != nil {
		        log.Println(n, err)
		        return
		    }
			fmt.Println(string(buf))
			out.WriteString(string(buf))
		}
	}()

	for {
		line, _, _ := in.ReadLine()
		
		n, err := conn.Write(line)
	    if err != nil {
	        log.Println(n, err)
	        return
    	}
	}
}