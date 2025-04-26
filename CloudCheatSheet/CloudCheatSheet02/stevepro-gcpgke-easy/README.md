# Cloud Setup Cheat Sheet GCP-GKE
dd-mmm-2025
<br />
Instructions for Cloud Setup Cheat Sheet blog post
<br />URL
<br /><br />
Pre-Requisites
```
gcloud auth login
gcloud auth application-default login
gcloud auth configure-docker
gcloud config set project SteveProProject
```
Check resources
```
gcloud compute instances list
gcloud compute disks list
gcloud compute forwarding-rules list
gcloud compute firewall-rules list
gcloud compute addresses list
gcloud container clusters list
```

Kubernetes [remote]
COMMAND #01 Create
```
gcloud container clusters create stevepro-gcp-gke           \
    --zone europe-west1-b                                   \
    --num-nodes 1                                           \
    --machine-type e2-medium                                \
    --image-type COS_CONTAINERD
```

COMMAND #02 Credentials
```
gcloud container clusters get-credentials stevepro-gcp-gke  \
    --zone=europe-west1-b                                   \
    --quiet --verbosity debug
```

COMMAND #03 DeployTest
```
kubectl create ns test-ns
kubectl config set-context --current --namespace=test-ns
kubectl apply -f Kubernetes.yaml
kubectl port-forward service/flask-api-service 8080:80
curl http://localhost:8080
```

COMMAND #04 Shell into Node
```
mkdir -p ~/GitHub/luksa
cd ~/GitHub/luksa
git clone https://github.com/luksa/kubectl-plugins.git
cd kubectl-plugins
chmod +x kubectl-ssh
kubectl get nodes
./kubectl-ssh node gke-stevepro-gcp-gke-default-pool-0b4ca8ca-sjpj
```

COMMAND #04 Destroy
```
gcloud container clusters delete stevepro-gcp-gke           \
    --zone europe-west1-b                                   \
    --quiet --verbosity debug
```