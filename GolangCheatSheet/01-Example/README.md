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
curl http://localhost:8081/test
```
```
02. Docker
Ctrl + Shift + P | Add Docker Files to Workspace
Go | 8081 | No
Right click Dockerfile | Build image... | 01example:latest
Right click 01example:latest | Run interactive
curl http://localhost:8081/test
```
```
03. Kubernetes
minikube start
minikube docker-env
eval $(minikube -p minikube docker-env)
docker build -t stevepro/testwebapi:1.0 .
kubectl apply -f Kubernetes.yaml
minikube service testwebapi-service --url
curl http://192.168.49.2:30799/test
kubectl delete -f Kubernetes.yaml
minikube stop
```
