package main

import (
	alexa "github.com/mikeflynn/go-alexa/skillserver"
	"log"
	"net/http"
)

var skills = map[string]interface{}{
	"/echo/computercontrol": alexa.EchoApplication{ // Route
		AppID:   "amzn1.ask.skill.058881dc-31d8-49a8-bf02-24684d82b9c0", // Echo App ID from Amazon Dashboard
		Handler: echoHandleIntent,
	},
}

func main() {
	alexa.Run(skills, "3000")
}

func echoHandleIntent(w http.ResponseWriter, r *http.Request) {
	request := alexa.GetEchoRequest(r)
	var response *alexa.EchoResponse

	if request.GetRequestType() == "IntentRequest" {
		switch request.GetIntentName() {
		case "lock":
			log.Println("Locking computer")
			response = alexa.NewEchoResponse().OutputSpeech("Computer has been locked.").Card("ComputerControl", "Computer has been locked.")
		}
	}

	json, _ := response.String()
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Write(json)
}
