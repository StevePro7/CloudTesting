Kubernetes ENV VARs
19-Apr-2023

01.
Define Environment Variables for a Container
https://kubernetes.io/docs/tasks/inject-data-application/define-environment-variable-container


kubectl create ns env-vars

kubectl config set-context --current --namespace=env-vars

01.
kubectl apply -f envars.yaml
kubectl exec envar-demo -- printenv

HOSTNAME=envar-demo
DEMO_FAREWELL=Such sweet sorrow
DEMO_GREETING=hELLO FROM THE ENVIRONMENT


kubectl get pods -l purpose=demo-envars
NAME         READY   STATUS    RESTARTS   AGE
envar-demo   1/1     Running   0          5m36s


02.
kubectl apply -f greeting.yaml
kubectl logs print-greeting
Warm greetings to The Most HOnorable Kuberentes

NB: I got CrashLoopBackOff initially
so  I deleted the resources section and re-ran OK
