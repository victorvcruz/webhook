package main

import (
	httpserver "webhooks-chat/http-server"
)

func init() {

}

func main() {

	api := httpserver.API{}

	api.Run()
}
