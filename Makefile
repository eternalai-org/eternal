# The binary to build, based on bin
BIN := eternal

test:

vendor:
	go mod tidy

api: vendor
	go build -o build/$(BIN) main.go

start_api: api
	./build/$(BIN) --config-file=config.env

clean:
	if [ -f ${BIN} ] ; then rm ${BIN} ; fi

lint-prepare:
	@echo "Installing golangci-lint"
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

docker:
	docker-compose down	service_miner && docker-compose build service_miner && docker-compose up -d service_miner &

docker-ollama:
	docker-compose down	 ollama && docker-compose build   ollama && docker-compose up -d   ollama &

abi-gen:
	abigen --pkg worker_hub --abi ./chains/base/contract/worker_hub/worker_hub.json --out ./chains/base/contract/worker_hub/worker_hub.go
	abigen --pkg erc20 --abi ./chains/base/contract/erc20/erc20.json --out ./chains/base/contract/erc20/erc20.go
	abigen --pkg staking_hub --abi ./chains/base/contract/staking_hub/staking_hub.json --out ./chains/base/contract/staking_hub/staking_hub.go

lint: vendor
	./bin/golangci-lint run ./... --timeout 10m0s

.PHONY: clean build test vendor lint-prepare lint
