/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"rabbitmq-messenger/lib"

	"github.com/spf13/cobra"
)

// chatCmd represents the chat command
var chatCmd = &cobra.Command{
	Use:     "chat ",
	Short:   "command for connecting to an arbitrary chat room with a nickname. if room is not exist, it will be created then",
	Example: `msgr chat --room <room_name> --nickname <nickname>`,

	Run: func(cmd *cobra.Command, args []string) {

		lib.StartChat(Room, Nickname)
	},
}
var Room string
var Nickname string

func init() {
	rootCmd.AddCommand(chatCmd)
	chatCmd.Flags().StringVarP(&Room, "room", "r", "", "please input the chatroom name/id")
	chatCmd.MarkFlagRequired("room")
	chatCmd.Flags().StringVarP(&Nickname, "nickname", "n", "", "please input the nickname for chatroom")
	chatCmd.MarkFlagRequired("nickname")

}
