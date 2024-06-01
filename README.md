# nginx-ingress-auto-scale

kubectl apply -f kubernetes/deployment.yaml
kubectl apply -f kubernetes/service.yaml

helm repo add ingress-nginx <https://kubernetes.github.io/ingress-nginx>
helm repo update
helm install nginx-ingress ingress-nginx/ingress-nginx
minikube addons enable metrics-server
kubectl apply -f <https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml>
kubectl get deployment metrics-server -n kube-system

kubectl apply -f kubernetes/ingress.yaml

minikube addons enable ingress
minikube tunnel
OR forward port with the nginx ingress service


cd test
go run test.go