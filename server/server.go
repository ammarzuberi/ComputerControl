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
	req := alexa.GetEchoRequest(r)

	switch req.GetIntentName() {
	case "lock":
		log.Println("Locking computer")
	}
}

func EchoIntentHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	echoResp.OutputSpeech("Hello world from my new Echo test app!").Card("Hello World", "This is a test card.")
}
