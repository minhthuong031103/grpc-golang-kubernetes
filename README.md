# grpc-golang-kubernetes

kubectl apply -f kubernetes/deployment.yaml kubectl apply -f kubernetes/service.yaml

helm repo add ingress-nginx <https://kubernetes.github.io/ingress-nginx> helm repo update helm install nginx-ingress ingress-nginx/ingress-nginx minikube addons enable metrics-server kubectl apply -f <https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml> kubectl get deployment metrics-server -n kube-system

kubectl apply -f kubernetes/ingress.yaml

minikube addons enable ingress minikube tunnel OR forward port with the nginx ingress service

cd test go run test.go

=> logic code order, proto common, performance testing

helm repo add prometheus-community <https://prometheus-community.github.io/helm-charts> helm repo add grafana <https://grafana.github.io/helm-charts> helm repo update helm install prometheus prometheus-community/prometheus helm install grafana grafana/grafana helm install prometheus-operator prometheus-community/kube-prometheus-stack

kubectl get secret --namespace default grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo

//using window powershell $encodedUsername = kubectl get secret --namespace default grafana -o jsonpath="{.data.admin-user}"

$encodedPassword = kubectl get secret --namespace default grafana -o jsonpath="{.data.admin-password}"

$decodedUsername = [System.Text.Encoding]::UTF8.GetString([System.Convert]::FromBase64String($encodedUsername))

$decodedPassword = [System.Text.Encoding]::UTF8.GetString([System.Convert]::FromBase64String($encodedPassword))

Write-Output "Username: $decodedUsername"

Write-Output "Password: $decodedPassword"

kubectl create -f prometheus-operator-crd kubectl apply -R -f .\monitoring\
