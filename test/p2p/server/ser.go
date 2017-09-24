package main

import (
	"encoding/binary"
	"encoding/hex"
	"log"
	"net"
	"os"
	"time"
)

const (
	CMD_LOGIN byte = byte(iota)
	CMD_LOGIN_RES

	CMD_LIST
	CMD_LIST_RES

	CMD_PING
	CMD_PONG

	CMD_CONE
	CMD_CONE_RES

	CMD_MEG
	CMD_MSG_RES
)

var userlist []*net.UDPAddr
var serverAddr *net.UDPAddr
var socket *net.UDPConn

func main() {
	args := os.Args

	if len(args) < 1 {
		log.Println("xxxx")
	}

	var err error

	log.SetFlags(log.Lshortfile)

	userlist = make([]*net.UDPAddr, 0, 10)

	serverAddr, err = net.ResolveUDPAddr("udp4", "192.168.30.200:8080")
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("server: ", serverAddr)

	socket, err = net.ListenUDP("udp4", serverAddr)
	if err != nil {
		log.Println("listen failed: ", err)
		return
	}
	defer socket.Close()

	log.Println("Listen...")

	for {
		data := make([]byte, 4096)
		read, addr, err := socket.ReadFromUDP(data)
		if err != nil {
			log.Println("readfromudp failed: ", err)
			continue
		}
		log.Printf("UDP: %d, %s, %s\n", read, addr, hex.EncodeToString(data[:read]))

		switch data[0] {
		case CMD_LOGIN:
			touser_list_data := make([]byte, 0, 15)
			touser_list_data = append(touser_list_data, CMD_LIST_RES, 0, 0, 0, 0, 0, 0)
			copy(touser_list_data[1:5], addr.IP)
			binary.LittleEndian.PutUint16(touser_list_data[5:], uint16(addr.Port))

			log.Println("touser_list_data:", hex.EncodeToString(touser_list_data))

			user_list_data := make([]byte, 0, 100)
			user_list_data = append(user_list_data, CMD_LIST_RES)

			for _, touser := range userlist {
				user_list_data = append(user_list_data, 0, 0, 0, 0, 0, 0)
				copy(user_list_data[len(user_list_data)-6:], touser.IP)
				binary.LittleEndian.PutUint16(user_list_data[len(user_list_data)-2:], uint16(touser.Port))

				socket.WriteToUDP(touser_list_data, touser)
			}

			log.Println("user_list_data:", hex.EncodeToString(user_list_data))
			socket.WriteToUDP(user_list_data, addr)

			userlist = append(userlist, addr)
		case CMD_LOGIN_RES:
		case CMD_LIST:
		case CMD_LIST_RES:

		case CMD_PING:
		case CMD_PONG:
			log.Println("CMD_PONG udp: ", addr)
		case CMD_CONE:
		case CMD_CONE_RES:
		case CMD_MEG:
		case CMD_MSG_RES:
		default:
			log.Println("default udp: ", addr)
		}
	}
}

func ping() {
	ping_data := make([]byte, 0, 15)
	ping_data = append(ping_data, CMD_PING, 0)

	for {
		for _, touser := range userlist {
			socket.WriteToUDP(ping_data, touser)
		}

		time.Sleep(5 * time.Second)
	}
}
