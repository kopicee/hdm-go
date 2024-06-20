default: help

up: ## Start the app
	docker compose up

test:  ## Run unit tests
	go install gotest.tools/gotestsum@v1.8.1
	gotestsum -f dots -- -failfast -covermode=count -coverprofile coverage.out ./...
	@go tool cover -func=coverage.out | grep 'total' | sed -e 's/\t\+/ /g'

help:  ## Show this help
	@echo 'usage: make [target] ...'
	@echo ''
	@echo 'targets:'
	@egrep '^(.+)\:\ .*##\ (.+)' ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'
