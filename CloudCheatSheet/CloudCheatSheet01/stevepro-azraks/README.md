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

Master SSH Key
```
cd ~/.ssh
ssh-keygen -t rsa -b 4096 -N '' -f master_ssh_key
eval $(ssh-agent -s)
ssh-add master_ssh_key
```
Kubernetes [remote]
```
az login
az group create --name stevepro-azraks-rg --location northeurope --debug
```

# 00	security principal
```
az ad sp create-for-rbac --name ${USER}-sp --skip-assignment
```

# 01	OUTPUT
```
{
    "appId": "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX",
    "displayName": "stevepro-sp",
    "name": "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX",
    "password": "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX",
    "tenant": "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX",
}
```
# 01	EXPORT
```
export AZ_SP_ID=<value_from_appId>
export AZ_SP_PASSWORD=<value_from_password>
```

# 02	create cluster
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

# 04	get credentials
```
export KUBECONFIG=~/.kube/config
az aks get-credentials --name stevepro-azraks   \
	--resource-group stevepro-azraks-rg --file ${KUBECONFIG}
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

COMMAND #05 Cleanup
```
kubectl delete -f Kubernetes.yaml
kubectl delete ns test-ns
```

# 05	delete cluster
```
az aks delete --name stevepro-azraks            \
    --resource-group stevepro-azraks
```