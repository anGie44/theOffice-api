check_install:
	which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_install
	swagger generate spec -o ./docs/swagger.yaml --scan-models

install:
	go mod tidy

run: install
	go run main.go

.PHONY: check_install swagger install run