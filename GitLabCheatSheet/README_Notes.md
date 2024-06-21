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
docker build --pull --rm -f "Dockerfile" -t flask-api:1.0 "."
docker run --rm -d -p 8080:8080/tcp flask-api:1.0 --name "flask-api"
curl http://localhost:8080
docker stop <conteinar_id>
```
```
03. Kubernetes [local]
kind create cluster --name flask-cluster
kubectl create ns test-ns
kubectl config set-context --current --namespace=test-ns
kind load docker-image flask-api:1.0 --name flask-cluster
kubectl apply -f Kubernetes.yaml
kubectl port-forward service/flask-api-service 8080:80
curl http://localhost:8080
kubectl delete -f Kubernetes.yaml
kind delete cluster --name flask-cluster
```
```
03. Kubernetes [remote]
```