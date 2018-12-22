package main

import (
	"fmt"
	"io/ioutil" 
	db "./db"
	bot "./bot"
)



func main() {
	db.SetupDB() 
	token, err := ioutil.ReadFile("token") 

	if err != nil {
		fmt.Println(err)
	}

	bot := bot.New(string(token))
	
	if bot != nil {
		bot.Run()
	}

}
