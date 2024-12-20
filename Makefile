# The binary to build, based on bin
BIN := eternal

test:

vendor:
	go mod tidy

api: vendor
	go build -o build/$(BIN) main.go

start_api: api
	./build/$(BIN) -lighthouse=12324242424 -chain=45762 -no-update-on-start=true -account_priv=6a6dd464ffabf8c6678324993032bca5ae4c4e92042e3900e68793cb911275f4 -api-key=ollama -api-url="https://localhost:14345/"

clean:
	if [ -f ${BIN} ] ; then rm ${BIN} ; fi

lint-prepare:
	@echo "Installing golangci-lint"
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

docker:
	docker-compose down	service_miner && docker-compose build service_miner && docker-compose up -d service_miner &

docker-ollama:
	docker-compose down	 ollama && docker-compose build   ollama && docker-compose up -d   ollama &

lint: vendor
	./bin/golangci-lint run ./... --timeout 10m0s

.PHONY: clean build test vendor lint-prepare lint
