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
az group create --location northeurope --name stevepro-azraks-rg 
```

# 02	create cluster
```
az aks create                                   \
    --resource-group stevepro-azraks-rg         \
    --name stevepro-azraks                      \
    --node-count 1                              \
    --ssh-key-value ~/.ssh/master_ssh_key.pub
```

# 03	delete cluster
```
az aks delete --name stevepro-azraks            \
    --resource-group stevepro-azraks
```




