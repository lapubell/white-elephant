package main

import "github.com/lapubell/white-elephant/server"

func main() {
	s, err := server.NewServer("9090")
	if err != nil {
		panic(err)
	}

	err = s.Serve()
	if err != nil {
		panic(err)
	}
}
