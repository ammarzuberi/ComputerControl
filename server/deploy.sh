GOOS=linux GOARCH=amd64 go build -v server.go && git add . && git commit -m "Deploy to development server" && git push && ssh -i /Users/az/Documents/Projects/Progerse/MyAWS.pem admin@ec2-54-187-129-24.us-west-2.compute.amazonaws.com 'cd /var/www/computercontrol.progerse.com/server && sudo git pull && sudo killall -9 server ; sudo screen -d -m /var/www/computercontrol.progerse.com/server/server'
