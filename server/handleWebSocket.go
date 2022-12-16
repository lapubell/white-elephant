package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func (s *server) handleWebSocket() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Configure the upgrader
		var upgrader = websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
		// Upgrade initial GET request to a websocket
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal(err)
		}
		// Make sure we close the connection when the function returns
		defer ws.Close()

		fmt.Println("adding person to ws pool")
		s.clients[ws] = Person{}

		for {
			var msg Message
			err := ws.ReadJSON(&msg)
			if err != nil {
				fmt.Println(err)
				delete(s.clients, ws)
				break
			}

			if msg.Type == "SET-NAME" {
				person := s.clients[ws]
				person.Name = msg.Message.(string)
				s.clients[ws] = person
				continue
			}

			if msg.Type == "SET-GIFT" {
				person := s.clients[ws]
				person.Gift = msg.Message.(string)
				s.clients[ws] = person
				continue
			}

			if msg.Type == "STEAL" {
				for i, p := range s.clients {
					if p.ChosenGift == msg.Message.(string) {
						p.ChosenGift = ""
						s.clients[i] = p
						break
					}
				}

				person := s.clients[ws]
				person.ChosenGift = msg.Message.(string)
				s.clients[ws] = person

				s.messageChannel <- s.messageAvailableGifts()

				continue
			}

			if msg.Type == "UNWRAP" {
				person := s.clients[ws]
				person.ChosenGift = msg.Message.(string)
				s.clients[ws] = person

				s.messageChannel <- s.messageAvailableGifts()

				continue
			}

			s.messageChannel <- msg
		}
	}
}
