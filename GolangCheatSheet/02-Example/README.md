##### README.md
###### xx/xx/2021
```
01. Locally
o mod init testwebapi
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
Right click Dockerfile | Build image... | 02example:latest
Right click 01example:latest | Run interactive
curl http://localhost:8081/test
```
```
03. Kubernetes
kind create cluster
docker build -t stevepro/testwebapi:2.0 .
kind load docker-image stevepro/testwebapi:2.0
kubectl apply -f Kubernetes.yaml
kubectl get nodes -o wide
kubectl get services
curl http://172.18.0.2:31196/test
kubectl delete -f Kubernetes.yaml
kind delete cluster
```
