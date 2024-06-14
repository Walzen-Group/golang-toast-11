package main

import (
	"../toast"
	"log"
	"os"
)


func main() {
	base, _ := os.Getwd()
	sonic := base + "\\examples\\sonic.jpg"

	notification := toast.Notification{
		AppID: "Example App",
		Title: "My notification",
		Message: "Some message about how important something is...",
		Icon: sonic,  // This file must exist (remove this line if it doesn't)
		Actions: []toast.Action{
			{"protocol", "Button 2", "", "", ""},
			{"protocol", "Button 2", "", "", ""},
		},
		Attribution: "Via Goland !!",
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}
