06-Dec-2024

kind create cluster

pip install flask

pip freeze > requirements.txt

F5

curl http://0.0.0.0:5000/
curl http://0.0.0.0:5000/read

curl --location --request POST 'http://0.0.0.0:5000/write' \
--header 'Content-Type: application/json' \
--data-raw '{
    "content": "This is the initial test"
}'


DOCKER
pyctest:latest
docker build --pull --rm -f "Dockerfile" -t pvctest:latest "." 


persistent-volume.yaml
persistent-volume-claim.yaml

deployment.yaml
service.yaml


kubectl apply -f persistent-volume.yaml
kubectl apply -f persistent-volume-claim.yaml
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml


kubectl port-forward service/flask-service 5000:80
curl http://0.0.0.0/


Shell into pod
cd /data
no files


curl http://0.0.0.0/
curl http://0.0.0.0/read

curl --location --request POST 'http://0.0.0.0/write' \
--header 'Content-Type: application/json' \
--data-raw '{
    "content": "This is the initial test"
}'

Shell into pod
cd /data
cat default.txt


Now delete the pod
kubectl delete po <pod_id>

Shell into pod
cd /data
cat default.txt


kind delete cluster