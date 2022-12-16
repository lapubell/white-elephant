package server

import "time"

func (s *server) broadcastPeople() {
	for {
		s.messageChannel <- Message{
			From:    "system",
			Type:    "people",
			Message: s.peopleList(),
		}
		time.Sleep(2 * time.Second)
	}
}

func (s *server) broadcastAvailableGifts() {
	for {
		s.messageChannel <- s.messageAvailableGifts()

		time.Sleep(5 * time.Second)
	}
}

func (s *server) messageAvailableGifts() Message {
	return Message{
		From:    "system",
		Type:    "available-gifts",
		Message: s.availableGiftList(),
	}
}
