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