# The binary to build, based on bin
BIN := eternal

test:

vendor:
	go mod tidy

app: vendor
	go build -o build/$(BIN) cmd/main.go

start_worker: app
	./build/$(BIN) --config-file=env/development.worker.yml

clean:
	if [ -f ${BIN} ] ; then rm ${BIN} ; fi

lint-prepare:
	@echo "Installing golangci-lint"
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

swagger:
	swag init --parseDependency --parseInternal -g cmd/main.go

lint: vendor
	./bin/golangci-lint run ./... --timeout 10m0s

abigen-prepare:
	go install github.com/ethereum/go-ethereum/cmd/abigen@latest

abi-gen:
	abigen --pkg erc20 --abi ./internal/contract/erc20/erc20.json --out ./internal/contract/erc20/erc20.go

.PHONY: clean build test vendor lint-prepare lint
