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
# 01	cluster YAML
create ~/stevepro-awseks/cluster.yaml
```
kind: ClusterConfig
apiVersion: eksctl.io/v1alpha5
metadata:
  name: stevepro-aws-eks
  region: eu-west-1
nodeGroups:
  - name: stevepro-aws-eks
    instanceType: m5.large
    desiredCapacity: 1
cloudWatch:
    clusterLogging:
      enableTypes: ["all"]
```

# 02	create cluster
```
eksctl create cluster -f ~/stevepro-awseks/cluster.yaml   \
    --kubeconfig ~/stevepro-awseks/kubeconfig             \
    --verbose 5
```

# 03	scale nodegroup
```
eksctl scale nodegroup                                    \
    --cluster=stevepro-aws-eks                            \
    --name=stevepro-aws-eks                               \
    --nodes=3                                             \
    --nodes-min=0                                         \
    --nodes-max=3                                         \
    --verbose 5
```

# 04 delete
kubectl delete -f Kubernetes.yaml
```
eksctl delete cluster                                     \
    --name=stevepro-aws-eks                               \
    --region eu-west-1                                    \
    --force			    
```