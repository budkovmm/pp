BINARY=ppserver
run_air:
	air -c .air.toml

build:
	go build -o ${BINARY} cmd/server/*.go

install-linter:
	@echo "Installing golangci-lint"
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

docker:
	docker build -t pp .

dc-run:
	docker-compose up --build -d

dc-stop:
	docker-compose down

lint:
	./bin/golangci-lint run ./...

unittest:
	go test -short  ./...