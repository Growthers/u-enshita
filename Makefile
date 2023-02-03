.PHONY: init
init:
	go mod download

.PHONY: start
start:
	go run main.go
