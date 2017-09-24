package main

import (
	"bufio"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
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
var listenAddr *net.UDPAddr
var socket *net.UDPConn

func main() {
	var err error

	log.SetFlags(log.Lshortfile)

	userlist = make([]*net.UDPAddr, 0, 10)

	serverAddr, err = net.ResolveUDPAddr("udp4", "192.168.30.200:8080")
	if err != nil {
		log.Println("serverAddr", err)
		return
	}

	log.Println("server: ", serverAddr)

	port := 8000
PORT:

	listenAddr, err = net.ResolveUDPAddr("udp4", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Println(err)
	}

	socket, err = net.ListenUDP("udp4", listenAddr)
	if err != nil {
		log.Println("ListenUDP:", err)
		port++
		goto PORT
		return
	}
	defer socket.Close()

	login_data := make([]byte, 0, 10)
	login_data = append(login_data, CMD_LOGIN)
	login_data = append(login_data, []byte("nickname")...)

	_, err = socket.WriteToUDP(login_data, serverAddr)
	if err != nil {
		log.Println("WriteToUDP: ", err)
		return
	}

	go readMsg()

	readCmd()
}

func readCmd() {
	for {
		fmt.Printf("p2p > ")

		scanner := bufio.NewScanner(os.Stdin)
		if !scanner.Scan() {
			continue
		}

		var line = scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Println("read error: ", err)
			continue
		}

		switch {
		case strings.HasPrefix(line, "help"):
			fmt.Println("  list: show all user list\n  send: send message\n\tsend <id> <message>")
		case strings.HasPrefix(line, "list"):
			fmt.Println("user list:")
			for id, user := range userlist {
				fmt.Println(id+1, user.IP, user.Port)
			}
		case strings.HasPrefix(line, "send"):
			id := 0
			content := ""
			fmt.Sscanf(line, "send %d %s", &id, &content)

			if id <= 0 || id > len(userlist) {
				fmt.Printf("error: id %d not fund\n", id)
				continue
			}

			log.Printf("send message: %s %d, %s", userlist[id-1], id, content)

			sendData := make([]byte, 0, 100)
			sendData = append(sendData, CMD_MEG)
			sendData = append(sendData, []byte(content)...)

			n, err := socket.WriteToUDP(sendData, userlist[id-1])
			if err != nil {
				log.Println(n, err)
			}
		case strings.HasPrefix(line, "quit"):
			return
		default:
			fmt.Printf("command error: %s\nuse the 'help' command to get help\n", line)
		}
	}
}

func readMsg() {
	for {
		data := make([]byte, 1024)
		read, addr, err := socket.ReadFromUDP(data)
		if err != nil {
			log.Println("ReadFromUDP: ", err)
			continue
		}
		log.Printf("UDP: %d, %s, %s\n", read, addr, hex.EncodeToString(data[:read]))

		switch data[0] {
		case CMD_LOGIN_RES:

		case CMD_LIST_RES:
			for i := 1; i < read; i += 6 {
				addrData := data[i:]
				touser := &net.UDPAddr{
					IP:   net.IP(addrData[:4]),
					Port: int(binary.LittleEndian.Uint16(addrData[4:])),
				}

				coneData := make([]byte, 0, 10)
				coneData = append(coneData, CMD_CONE)
				coneData = append(coneData, []byte("nickname")...)

				socket.WriteToUDP(coneData, touser)
				log.Println("cone: ", touser, coneData)

				userlist = append(userlist, touser)
			}
		case CMD_PING:
			log.Printf("CMD_PING\n")
			pong_data := make([]byte, 0, 15)
			pong_data = append(pong_data, CMD_PONG, 1)
			n, err := socket.WriteTo(pong_data, addr)
			log.Println("CMD_PING: ", n, err)
		case CMD_PONG:

		case CMD_CONE:
			coneResData := make([]byte, 0, 10)
			coneResData = append(coneResData, CMD_CONE_RES)
			coneResData = append(coneResData, []byte("nickname")...)

			socket.WriteToUDP(coneResData, addr)
		case CMD_CONE_RES:
			log.Println("CMD_CONE_RES:", addr)
		case CMD_MEG:
			fmt.Println(string(data[1:read]))
		case CMD_MSG_RES:

		default:
			log.Println("default UDP: ", data[0])
		}
	}
}
