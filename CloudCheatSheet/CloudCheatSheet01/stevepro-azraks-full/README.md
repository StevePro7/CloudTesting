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
az ad sp create-for-rbac --name stevepro-sp --skip-assignment
```

# 02	create cluster
```
az aks create --name stevepro-azraks            \
    --resource-group stevepro-azraks            \
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

# 03	delete cluster
```
az aks delete --name stevepro-azraks            \
    --resource-group stevepro-azraks
```