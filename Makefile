default: help

up: ## Start the app using Docker Compose
	docker compose up --remove-orphans --build --detach

run: ## Start the app using native Go installation
	go build -o hdm-go
	./hdm-go

test:  ## Run unit tests
	go install gotest.tools/gotestsum@v1.8.1
	gotestsum -f dots -- -failfast -covermode=count -coverprofile coverage.out ./...
	@go tool cover -func=coverage.out | grep 'total' | sed -e 's/\t\+/ /g'

help:  ## Show this help
	@echo 'usage: make [target] ...'
	@echo ''
	@echo 'targets:'
	@egrep '^(.+)\:\ .*##\ (.+)' ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'
