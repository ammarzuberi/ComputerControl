package main

import (
	"fmt"
	"github.com/go-redis/redis"
	alexa "github.com/mikeflynn/go-alexa/skillserver"
	"log"
	"net/http"
)

var redisClient *redis.Client

var applications = map[string]interface{}{
	"/echo/computercontrol": alexa.EchoApplication{
		AppID:   "amzn1.ask.skill.058881dc-31d8-49a8-bf02-24684d82b9c0",
		Handler: echoHandleIntent,
	},
	"/getCommand": alexa.StdApplication{
		Handler: handleCommandCallback,
		Methods: "GET",
	},
}

func handleCommandCallback(w http.ResponseWriter, r *http.Request) {
	currentCommand, err := redisClient.Get("command").Result()
	if err != nil {
		fmt.Fprintf(w, "Error")
	} else {
		fmt.Fprintf(w, currentCommand)
	}
}

func echoHandleIntent(w http.ResponseWriter, r *http.Request) {
	request := alexa.GetEchoRequest(r)
	var response *alexa.EchoResponse

	if request.GetRequestType() == "IntentRequest" {
		switch request.GetIntentName() {
		case "lock":
			log.Println("Locking computer")

			if value, _ := redisClient.Get("command").Result(); value == "wait" {
				redisClient.Set("command", "lock", 0)
			}

			response = alexa.NewEchoResponse().OutputSpeech("Computer has been locked.").Card("ComputerControl", "Computer has been locked.")
		}

		//Add new intents here by name
	}

	json, _ := response.String()
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Write(json)
}

func main() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatal("Could not connect to Redis database")
	}
	redisClient.Set("command", "wait", 0)

	alexa.Run(applications, "3000")
}
