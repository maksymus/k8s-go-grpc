.PHONY: proto
proto:
	protoc -I ./proto/ --go_out=plugins=grpc:./proto/ ./proto/*.proto

.PHONY: gcd
gcd:
	go build -o gcd-service/gcd gcd-service/main.go

.PHONY: api
api:
	go build -o api-service/api api-service/main.go

compile: gcd api

build: compile
	docker build -t local/gcd -f gcd-service/Dockerfile .
	docker build -t local/api -f api-service/Dockerfile .

# kubectl get pods -w - to check if deployed
deploy:
	kubectl create -f gcd-service/deployment.yaml
	kubectl create -f api-service/deployment.yaml

undeploy:
	kubectl delete -f gcd-service/deployment.yaml
	kubectl delete -f api-service/deployment.yaml

redeploy: undeploy deploy