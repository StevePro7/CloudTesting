Kubernetes ENV VARs
19-Apr-2023

01.
Define Environment Variables for a Container
https://kubernetes.io/docs/tasks/inject-data-application/define-environment-variable-container


kubectl create ns env-vars

kubectl config set-context --current --namespace=env-vars


kubectl apply -f envars.yaml
kubectl exec envar-demo -- printenv

HOSTNAME=envar-demo
DEMO_FAREWELL=Such sweet sorrow
DEMO_GREETING=hELLO FROM THE ENVIRONMENT


kubectl get pods -l purpose=demo-envars
NAME         READY   STATUS    RESTARTS   AGE
envar-demo   1/1     Running   0          5m36s
