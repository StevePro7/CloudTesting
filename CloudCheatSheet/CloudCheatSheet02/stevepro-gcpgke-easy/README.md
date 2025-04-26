# Cloud Setup Cheat Sheet
dd-mmm-2025
<br />
Instructions for Cloud Setup Cheat Sheet blog post
<br />URL
<br /><br />
Pre-Requisites
```
TODO
```

Kubernetes [remote]
Pre-Requisites
```
gcloud auth login
gcloud auth application-default login
gcloud auth configure-docker
gcloud config set project SteveProProject
```
Check resources
```
gcloud compute regions list
gcloud compute zones list
gcloud compute instances list
gcloud compute disks list
gcloud compute forwarding-rules list
gcloud compute firewall-rules list
gcloud compute addresses list
gcloud container clusters list
```
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

COMMAND #03 Destroy
```
gcloud container clusters delete stevepro-gcp-gke           \
    --zone europe-west1-b                                   \
    --async --quiet --verbosity debug
```