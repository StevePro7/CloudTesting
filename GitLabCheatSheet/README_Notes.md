##### README.md
###### 04-Jul-2024
```
01. Locally
pip install -r requirements.txt
pip install --upgrade pip
touch main.py
python main.py
curl http://localhost:8080
```
```
02. Docker
Ctrl + Shift + P | Add Docker Files to Workspace
Go | 8081 | No
Right click Dockerfile | Build image... | 02example:latest
Right click 02example:latest | Run interactive
curl http://localhost:8081/test
```
```
03. Kubernetes [local]
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
```
03. Kubernetes [remote]
```