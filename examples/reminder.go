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
		ActivationArguments: "action=viewEvent&amp;eventId=1983",
		Actions: []toast.Action{
			{ Type: "system", Arguments: "snooze", HintInputId: "snoozeTime" },
			{ Type: "system", Arguments: "dismiss"},
		},
		Input: toast.Input{
			Id:         "snoozeTime",
			Type:       "selection",
			DefaultInput: "15",
			Selections: []toast.Selection{
				{"1", "1 minute"},
				{"15", "15 minutes"},
				{"60", "1 hour"},
				{"240", "4 hours"},
				{"1140", "1 day"},
			},
		},
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}

}
