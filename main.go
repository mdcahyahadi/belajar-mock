package main

import "httpauth/helpers"

var PORT = ":8080"

func main() {
	helpers.StartServer().Run(PORT)
}
