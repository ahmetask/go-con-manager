docker build -t go-con-manager .
kubectl delete deployment go-con-manager
kubectl delete service go-con-manager
kubectl create -f deployment.yml
kubectl create -f service.yml
