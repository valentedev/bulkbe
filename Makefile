.PHONY: run/app
run/app:
	@echo 'Starting server...'
	go run ./cmd/app

.PHONY: audit
audit:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...

.PHONY: vendor
vendor:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies...'
	go mod vendor

.PHONY: build/app
build/app:
	@echo 'Building cmd/api...'
	go build -ldflags="-s" -o=./bin/app ./cmd/app
	GOOS=linux GOARCH=amd64 go build -ldflags="-s" -o=./bin/linux_amd64/app ./cmd/app