minikube start

eval $(minikube -p minikube docker-env)
docker build . -t my-go-app
docker pull busybox

helm repo add istio https://istio-release.storage.googleapis.com/charts
helm repo update
helm install istio-base istio/base -n istio-system --set defaultRevision=default --create-namespace
helm install istiod istio/istiod -n istio-system --wait
helm install istio-ingress istio/gateway --wait # !!! Open another terminal and enter command: minikube tunnel

kubectl apply -f yamls/