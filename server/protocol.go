//
// Copyright 2014 Hong Miao. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	//"log"
	"fmt"
	"errors"
)

type CmdFuc func(server *SecureServer, client *SecureClient, args []string)

var (
	commands map[string]CmdFuc
)

func init() {
	commands = map[string]CmdFuc {
		"/setName" : setName,
		"/weather" : weather,
		"/ticket"  : ticket,
	}
}

type Cmd struct {
	cmd string
	args []string
}

func CreateCmd() *Cmd {
	return &Cmd {
		
	}
}

func (self *Cmd)parseCmd(msglist []string) {
	if len(msglist) == 1 {
		self.cmd = msglist[0]
	} else {
		self.cmd = msglist[0]
		self.args = msglist[1:]
	}
}

func (self *Cmd) executeCommand(server *SecureServer, client *SecureClient) (err error) {
	//fmt.Println("executeCommand")
	if f, ok := commands[self.cmd]; ok {
		f(server, client, self.args)
		return
	}

	err = errors.New("Unsupported command: " + self.cmd)
	return
}

func checkUsername(server *SecureServer, name string) bool {
	for _, client := range server.clients {
		if name == client.GetName() {
			return false
		}
	}
	
	return true
}

func setName(server *SecureServer, client *SecureClient, args []string) {
	//fmt.Println("setName")
	oldname := client.GetName()
	
	flag := checkUsername(server, args[0])
	//fmt.Println(flag)
	if flag {
		client.SetName(args[0])
		server.messageProcessRaw(fmt.Sprintf("Notification: %s changed name to %s", oldname, args[0]))
		server.messageProcessRaw(fmt.Sprintf("Welcome  %s to join the ChatRoom \n" ,args[0]))
		server.messageProcessRaw(fmt.Sprintf("Current have %d people on line \n", server.getOnline()))
	} else {
		server.sendToSingleClient(client, fmt.Sprintf("username is already used, Please input your username : \n"))
	}
}

func weather(server *SecureServer, client *SecureClient, args []string) {
	addr := args[0]
	server.messageProcessRaw(fmt.Sprintf("%s is sunnshine", addr))
}

func ticket(server *SecureServer, client *SecureClient, args []string) {
	from := args[0]
	to := args[1]
	server.messageProcessRaw(fmt.Sprintf("from %s to %s is sell out", from, to))
}

