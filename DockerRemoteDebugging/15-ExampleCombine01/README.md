##### README.md
###### 15/11/2022

Reference: https://medium.com/lightbaseio/web-application-firewall-in-go-feat-owasp-modsecurity-core-rule-set-3f97a26e3311
```
01. Locally
go mod init testwebapi
go get "github.com/gorilla/mux"
go mod tidy
touch main.go
go run main.go
curl http://localhost:8081/
```
```
02. Docker
Create Dockerfile
docker build -t example:latest .
docker run --rm -it -p 8081:8081/tcp example:latest
curl http://localhost:8081/


Bind socket already in use
sudo fuser -k 8081/tcp


NEW
Repeat 10-Example
but update Dockerfile

01.
FROM golang:latest

02.
#COPY . .
COPY . /app

03.
EXPOSE 8081 40000
Edit configurations
+
Go Remote
change port to 40000
set breakpoint
F5
curl localhost:8081
does not break
curl localhost:40000
connection refused

04.
# Installing Delve
RUN CGO_ENABLED=1 go install -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv@latest
F5
curl localhost:8081
does not break

05. interesting!
ENV GO111MODULE=off
causes NOT to be able to find gmux package??

06.
RUN CGO_ENABLED=1 go build -gcflags "all=-N -l" -o main .
works as before 
no break on F5

07.
comment out entry point
docker build
docker run
confronted by command prompt
root@97495feb07c0:/app# 
can run main directly
e.g.
root@97495feb07c0:/app# ./main 
on another terminal
curl loaclhost:8081
"Healthy...'20'!!"
no break on F5

08.

