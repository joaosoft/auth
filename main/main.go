package main

import "github.com/joaosoft/auth"

func main() {
	m, err := auth.NewAuth()
	if err != nil {
		panic(err)
	}

	if err := m.Start(); err != nil {
		panic(err)
	}
}
