package main

import (
	"flag"
	"fmt"
	"github.com/nghnam/gorocket"
	"os"
)

var (
	url      string
	username string
	password string
	token    string
	userid   string
)

func getConfig() {
	url = os.Getenv("RC_URL")
	username = os.Getenv("RC_USER")
	password = os.Getenv("RC_PASS")
	token = os.Getenv("RC_TOKEN")
	userid = os.Getenv("RC_USERID")
}

func main() {
	getConfig()
	var cred gorocket.Credentical
	if token == "" || userid == "" {
		cred, _ = gorocket.GetToken(url, username, password)
	}
	client := gorocket.Client{Url: url, Token: cred.Token, UserId: cred.UserId}
	listRoom := flag.Bool("list-room", false, "list public rooms")
	getVersion := flag.Bool("get-version", false, "get api version")
	sendMessage := flag.Bool("send", false, "send message")
	roomID := flag.String("room-id", "", "room ID")
	roomName := flag.String("room-name", "", "room Name")
	message := flag.String("message", "", "message content")
	flag.Parse()
	if *listRoom && !*getVersion && !*sendMessage {
		rooms, _ := client.ListPublicRooms()
		for _, room := range rooms {
			fmt.Println(room.Name, room.ID)
			return
		}
	}
	if !*listRoom && *getVersion && !*sendMessage {
		version, _ := gorocket.GetVersion(url)
		fmt.Println("API version: ", version.Api)
		fmt.Println("RocketChat version: ", version.RocketChat)
		return
	}
	if !*listRoom && !*getVersion && *sendMessage {
		if *message == "" {
			fmt.Println("Message content not exists")
			return
		}
		if *roomName != "" && *roomID == "" {
			id, err := client.GetRoomId(*roomName)
			if err != nil {
				fmt.Println("Room not found")
				return
			}
			err = client.SendMessage(id, *message)
			if err != nil {
				fmt.Println("Message not send")
				return
			}
		} else {
			err := client.SendMessage(*roomID, *message)
			if err != nil {
				fmt.Println("Message not send")
				return
			}
		}
	}
}
