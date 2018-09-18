# ComputerControl
A small and feature-limited implementation of voice controls for your computer via Alexa. This was created primarily to make it possible for me to tell Alexa to lock my computer as I walked out of the room.

## Extending ComputerControl

### Server Side
ComputerControl is very easy to extend. Simply add a new intent through Amazon's Alexa Dashboard and set up what you'd like the server to do upon that intent firing in `server/server.go`. If you'd like to push a command to the client, simply update the `command` key stored in Redis.

### Client Side
The client will reach out to the server every 5 seconds and ask for the command in Redis. You can easily customize the code in `client/client.go` to do whatever you'd like for different commands.

#### Using intents with slots
ComputerControl does not currently support this. As I find more uses for this setup, there will likely be added support for slots. For purposes of just locking, slots were not required. It should be relatively easy to receive and process the slots using the skillserver library that is currently used.