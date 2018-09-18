package main

import (
	alexa "github.com/mikeflynn/go-alexa/skillserver"
)

var Applications = map[string]interface{}{
	"/echo/helloworld": alexa.EchoApplication{ // Route
		AppID:    "amzn1.ask.skill.058881dc-31d8-49a8-bf02-24684d82b9c0", // Echo App ID from Amazon Dashboard
		OnIntent: EchoIntentHandler,
		OnLaunch: EchoIntentHandler,
	},
}

func main() {
	alexa.Run(Applications, "3000")
}

func EchoIntentHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	echoResp.OutputSpeech("Hello world from my new Echo test app!").Card("Hello World", "This is a test card.")
}
