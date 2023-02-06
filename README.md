# proglog
proglog implementation from Distributed Services with Go book

# service per pod

# Apply all set of production resources defined in kustomization.yaml in `production` directory .
kubectl apply -k https://github.com/metacontroller/metacontroller/manifests/production


 helm install proglog proglog \
--set image.repository=gcr.io/proglog-peauliol/proglog \
--set service.lb=tru

go run cmd/getservers/main.go -addr=35.225.255.85:8400 
servers
go run cmd/getservers/main.go -addr=34.134.134.61:8400
servers
go run cmd/getservers/main.go -addr=35.238.72.19:8400 

