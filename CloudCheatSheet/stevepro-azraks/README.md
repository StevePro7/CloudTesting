## Azure AKS
dd-mmm-2025
<br />
Instructions for Cloud Setup Cheat Sheet blog post
<br />URL
<br /><br />

Microsoft provides Azure Kubernetes Services as fully managed Kubernetes container orchestration service.
<br />
Follow all instructions below in order to provision a Kubernetes cluster and end-to-end test its functionality. 

#### Master SSH Key
```
cd ~/.ssh
ssh-keygen -t rsa -b 4096 -N '' -f master_ssh_key
eval $(ssh-agent -s)
ssh-add master_ssh_key
```

#### Pre-Requisites
```
az login
```

#### Check Resources
```
az account list --output table
az group list --output table
az resource list --output table
az resource list --query "[?location=='northeurope']" --output table
az vm list --output table
az aks list --output table
az container list --output table
az storage account list --output table
az network public-ip list --output table
```

#### Resource Group
```
az group create --name stevepro-azraks-rg --location northeurope --debug
```

#### Security Principal
```
az ad sp create-for-rbac --name ${USER}-sp --skip-assignment
```

#### Output
```
{
    "appId": "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX",
    "displayName": "stevepro-sp",
    "name": "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX",
    "password": "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX",
    "tenant": "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX",
}
```

#### Export
```
export AZ_SP_ID=<value_from_appId>
export AZ_SP_PASSWORD=<value_from_password>
```

#### Create Cluster
```
az aks create --name stevepro-azraks            \
    --resource-group stevepro-azraks-rg         \
    --dns-name-prefix stevepro-azraks           \
    --node-count 3                              \
    --node-vm-size Standard_D2s_v3              \
    --kubernetes-version 1.28                   \
    --ssh-key-value ~/.ssh/master_ssh_key.pub   \
    --service-principal ${AZ_SP_ID}             \
    --client-secret ${AZ_SP_PASSWORD}           \
    --load-balancer-sku standard                \
    --network-plugin azure --debug
```

#### Get Credentials
```
export KUBECONFIG=~/.kube/config
az aks get-credentials --name stevepro-azraks   \
	--resource-group stevepro-azraks-rg --file ${KUBECONFIG}
```

#### Deploy Test
```
kubectl create ns test-ns
kubectl config set-context --current --namespace=test-ns
kubectl apply -f Kubernetes.yaml
kubectl port-forward service/flask-api-service 8080:80
curl http://localhost:8080
```

#### Output
```
Hello World (Python)!
```

#### Shell into Node
```
mkdir -p ~/GitHub/luksa
cd ~/GitHub/luksa
git clone https://github.com/luksa/kubectl-plugins.git
cd kubectl-plugins
chmod +x kubectl-ssh
kubectl get nodes
./kubectl-ssh node TODO-get-AKS-node
```

#### Cleanup
```
kubectl delete -f Kubernetes.yaml
kubectl delete ns test-ns
```

#### Delete Cluster
```
az aks delete --name stevepro-azraks            \
    --resource-group stevepro-azraks-rg
```