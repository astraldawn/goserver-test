package main

import "github.com/go-martini/martini"

func main() {
	server := martini.Classic()
	server.Get("/", func() string {
		return "<h1>Hello, world!</h1>"
	})

	server.Run()
}
