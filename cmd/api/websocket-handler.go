package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"
)

// Saving user IDs in a map to their websocket connections
// Saving group members' user ID's in a map
var userConns = make(map[string]*websocket.Conn)
var groupMembers = make(map[int64][]string)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func webSocketHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userID")

	log.Printf("[NEW] A new connection detected!")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection to WebSocket: ", err)
		return
	}
	defer conn.Close()

	userConns[userID] = conn
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error when reading message: ", err)
			break
		}

		var message Message
		err = json.Unmarshal(msg, &message)
		if err != nil {
			log.Println("Error unmarshaling JSON message: ", err)
			continue
		}

		// Convert RoomID to int64
		// authorID, err := strconv.ParseInt(message.AuthorID, 10, 64)
		// if err != nil {
		// 	log.Println("Error converting AuthorID to int64: ", err)
		// 	continue
		// }

		roomID, err := strconv.ParseInt(message.RoomID, 10, 64)
		if err != nil {
			log.Println("Error converting RoomID to int64: ", err)
			continue
		}
		addUserToRoom(roomID, userID)
		// Handle message broadcast to group
		broadcastMessage(roomID, msg, message.AuthorID)
	}
}

func addUserToRoom(roomID int64, userID string) {
	for _, id := range groupMembers[roomID] {
		if id == userID {
			return // User already in the room
		}
	}
	groupMembers[roomID] = append(groupMembers[roomID], userID)
	log.Printf("User %s added to room %d", userID, roomID)
}

func broadcastMessage(roomID int64, msg []byte, authorID string) {
	for _, userID := range groupMembers[roomID] {
		conn, ok := userConns[userID]
		if ok {
			log.Println(authorID)
			err := conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("Error when sending message: ", err)
			}
		}
	}
}
