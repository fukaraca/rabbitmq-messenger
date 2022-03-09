package lib

import (
	"bufio"
	"os"
	"os/exec"
	"rabbitmq-messenger/consumer"
	"rabbitmq-messenger/producer"
	"runtime"
	"strings"
)

func StartChat(room, nick string) {
	reader := bufio.NewReader(os.Stdin)

	chP := producer.Connect(room)
	defer chP.Close()

	chC := consumer.Connect(room)
	defer chC.Close()

	callClear()
	consumer.Listen(chC, room, nick)

	for {
		msg, _ := reader.ReadString('\n')
		msg = replacer(msg)
	}

}

//replaces line breaks according to OS
func replacer(text string) string {

	old := ""
	if runtime.GOOS == "linux" {
		old = "\n"
	} else if runtime.GOOS == "windows" {
		old = "\r\n"
	}
	text = strings.Replace(text, old, "", -1)
	text = strings.ToLower(text)
	return text
}

var clear map[string]func() //create a map for storing clear funcs
func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

//clear terminal
func callClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
