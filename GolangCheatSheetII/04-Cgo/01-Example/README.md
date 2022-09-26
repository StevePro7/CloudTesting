##### README.md
###### xx/xx/2022

Reference: https://medium.com/lightbaseio/web-application-firewall-in-go-feat-owasp-modsecurity-core-rule-set-3f97a26e3311
```
01. Locally
go mod init testwebapi
touch main.go
touch Dockerfile
touch Kuberenetes.yaml

mkdir rules
copy 3x rules into rules folder

mkdir scripts
touch scripts/modsec.sh
chmod +x scripts/modsec.sh
Reference https://github.com/projectcalico/go-build/blob/master/scripts/modsec.sh

mkdir waf
cd waf
touch waf.go waf.h waf.c

Launch Goland IDE
complete code

main.go
go get "github.com/gorilla/mux"
go mod tidy

go run main.go
curl http://localhost:8081/
curl http://localhost:8081/test/artists.php
curl http://localhost:8081/test/artists.php?artists.php?artist=0+div+1+union%23foo*%2F*bar%0D%0Aselect%23foo%0D%0A1%2C2%2Ccurrent_user

```
```
02. Docker
Ctrl + Shift + P | Add Docker Files to Workspace
Go | 8081 | No
Right click Dockerfile | Build image... | 01example:latest
OR
docker build --pull --rm -f "Dockerfile" -t 01example:latest "."

Right click 01example:latest | Run interactive
OR
docker run -d -p 8081:8081 01example:latest
curl http://localhost:8081/
curl http://localhost:8081/test/artists.php
curl http://localhost:8081/test/artists.php?artists.php?artist=0+div+1+union%23foo*%2F*bar%0D%0Aselect%23foo%0D%0A1%2C2%2Ccurrent_user

OR
docker exec -it [containerID] bash
```
```
03. Kubernetes
minikube start
minikube docker-env
eval $(minikube -p minikube docker-env)
docker build -t stevepro/testwebapi:3.0 .
kubectl apply -f Kubernetes.yaml
minikube service testwebapi-service --url
curl http://192.168.49.2:32743/
curl http://192.168.49.2:32743/test/artists.php
curl http://192.168.49.2:32743/test/artists.php?artists.php?artist=0+div+1+union%23foo*%2F*bar%0D%0Aselect%23foo%0D%0A1%2C2%2Ccurrent_user
kubectl logs -f testwebapi-787c48fd87-shc6k 
kubectl delete -f Kubernetes.yaml
minikube stop
```
