/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
msgr chat -room "room-identifier" -nickname "nick"
msgr message -to "username" -from "username2"
*/
package main

import "rabbitmq-messenger/cmd"

func main() {
	cmd.Execute()
}
