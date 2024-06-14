package main

import (
	"log"
	"os"
)

import "../toast"

func main() {
	base, _ := os.Getwd()
	sonic := base + "\\examples\\sonic.jpg"

	notification := toast.Notification{
		AppID: "Example App",
		Title: "My notification",
		Message: "Some message about how important something is...",
		Icon: sonic,  // This file must exist (remove this line if it doesn't)
		Header: toast.Header{
			Id:        "6289",
			Title:     "Camping!!",
			Arguments: "action=openConversation&amp;id=6289",
		},
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}
