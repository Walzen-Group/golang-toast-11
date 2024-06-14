package main

import (
	"github.com/Walzen-Group/golang-toast-11"
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
			{"protocol", "Button 1", "", "", ""},
			{"protocol", "Button 2", "", "", ""},
		},
		ProgressBar: toast.ProgressBar{
			Title:               "Weekly playlist",
			Value:               0.6,
			Status:              "Downloading...",
			ValueStringOverride: "15/26 songs",
		},
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}
