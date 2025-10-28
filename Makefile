.PHONY: build test lint docker deploy run clean

build:
	go build -o bin/operator ./cmd/operator

test:
	go test ./... -v

lint:
	golangci-lint run ./...

docker:
	docker build -t gigvault/operator:local .

deploy:
	kubectl apply -f manifests/

run:
	go run ./cmd/operator/main.go

clean:
	rm -rf bin/
	go clean

