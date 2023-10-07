package main

import (
	"flag"
	"log"
)

func must_token() string {
	token := flag.String("token-bot-token", "", "token for access to telegram bot" )
	flag.Parse()
	if *token == ""{
		log.Fatal("token is not specified")
	}
	return *token
}

func main(){

}

