## Build the image using the following command

```
➜  ./create_image.sh
```


## Run minikube
```
➜  minikube start  
➜  minikube addons enable ingress 

```

# Minikube ip
```
➜  minikube ip
```

- In `/etc/hosts` make entry like this replace 127.0.0.1 with minikube ip

```
127.0.0.1                 challenge.local
127.0.0.1       detectify.challenge.local
```

- Using Kustomize
```
➜  kubectl apply -k ./kustomize/base/

➜  curl challenge.local/
{"status":"Hello world!"}
➜ curl detectify.challenge.local/    
{"status":"Hello World again"}
```

- Delete everything
```
➜ kubectl delete -k ./kustomize/base/
service "service" deleted
deployment.apps "deployment" deleted
ingress.networking.k8s.io "hello-again-ingress" deleted
ingress.networking.k8s.io "hello-ingress" deleted

➜  minikube delete
```
