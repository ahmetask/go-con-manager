docker build -t go-con-manager .
kubectl delete deployment go-con-manager
kubectl delete service go-con-manager
kubectl apply -f deployment.yml
kubectl apply -f service.yml