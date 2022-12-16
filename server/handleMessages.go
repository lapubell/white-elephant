package server

import "fmt"

func (s *server) handleMessages() {
	fmt.Println("now handling all messages")
	for {
		msg := <-s.messageChannel

		for client := range s.clients {
			err := client.WriteJSON(msg)
			if err != nil {
				fmt.Println(err)
				client.Close()
				delete(s.clients, client)
			}
		}
	}
}
