package main

import (
	"github.com/everdev/mack"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func main() {
	for {
		request, err := http.Get("https://computercontrol.mydomain.com/callback/getCommand")
		if err != nil {
			alert := mack.AlertOptions{
				Title:   "ComputerControl Error",
				Message: "Could not connect to command server.",
				Style:   "critical",
			}
			mack.AlertBox(alert)

			os.Exit(-1)
		}
		defer request.Body.Close()

		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			alert := mack.AlertOptions{
				Title:   "ComputerControl Error",
				Message: "Could not retrieve data from command server.",
				Style:   "critical",
			}
			mack.AlertBox(alert)

			os.Exit(-1)
		}

		switch string(body) {
		case "lock":
			command := exec.Command("/System/Library/CoreServices/Menu Extras/User.menu/Contents/Resources/CGSession", "-suspend")
			command.Run()
		default:
			//Do nothing
		}

		time.Sleep(5 * time.Second)
	}
}
