env:
	docker-compose up -d

run:
	go run ./bin/launcher/main.go

build:
	docker build -t auth:1.0 .

push:
	docker login --username joaosoft
	docker tag auth:1.0 joaosoft/auth
	docker push joaosoft/auth

fmt:
	go fmt ./...

vet:
	go vet ./*

gometalinter:
	gometalinter ./*