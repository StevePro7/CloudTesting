# Cloud Setup Cheat Sheet
dd-mmm-2025
<br />
Instructions for Cloud Setup Cheat Sheet blog post
<br />URL
<br /><br />
Pre-Requisites
```
```

Kubernetes [remote]
```
#az group create --name ${CLUSTER_NAME} --location ${AZ_LOCATION} --debug
```

# 00	security principal
```
az login
az ad sp create-for-rbac --name ${USER}-sp --skip-assignment
```

# 02	create cluster
```
az aks create --name stevepro-azraks			\
    --resource-group stevepro-azraks			\
    --dns-name-prefix stevepro-azraks			\
    --node-count 3								\
    --node-vm-size Standard_D2s_v3				\
    --kubernetes-version 1.28					\
    --ssh-key-value ~/.ssh/master_ssh_key.pub	\
    --service-principal ${AZ_SP_ID}				\
    --client-secret ${AZ_SP_PASSWORD}			\
    --load-balancer-sku standard				\
    --network-plugin azure --debug
```

# 03	delete cluster
```
az aks delete --name stevepro-azraks			\
    --resource-group stevepro-azraks
```