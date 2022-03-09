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
	Short:   "command for connecting to a certain chat room with a nickname",
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
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
