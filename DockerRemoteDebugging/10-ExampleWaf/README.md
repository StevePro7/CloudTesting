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
docker build -t 02example:latest
docker run --rm -it -p 8081:8081/tcp 02example:latest
curl http://localhost:8081/
