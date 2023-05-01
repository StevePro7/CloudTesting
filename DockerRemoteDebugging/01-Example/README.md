Remote Debugging
29-Mar-2023

https://kupczynski.info/posts/remote-debug-go-code
https://github.com/igor-kupczynski/remote-debug-example

go 1.18


got some weird issue about 1.14 too old so had to update the commands
# RUN go install github.com/go-delve/delve/cmd/dlv
RUN go install github.com/go-delve/delve/cmd/dlv@2f13672765fe


Run program directly
curl localhost:8080

copy in Docker files

Weird first time around
docker-compose up

error complaining about go version??
dropped to 1.14 - rebuild - OK then bump up and re-build OK again

curl localhost:8080
Hello world

but 8081 Connection refused

Update configuration


REPEAT
docker-compose up
curl localhost:8080
Hello World

curl localhost_8081
curl: (7) Failed to connect to localhost port 8081: Connection refused

Goland | Edit configurations
Go Remote
change port to 40000
set breakpoint in main.go
e.g.
now := time.Now()

F5
curl localhost_8081
code breaks at break point and can debug step thru code!