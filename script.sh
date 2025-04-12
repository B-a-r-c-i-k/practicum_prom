minikube start
eval $(minikube -p minikube docker-env)
docker build . -t my-go-app
docker pull busybox
kubectl apply -f yamls/