APP = vendingServer

.PHONY: build
build:
	go build -o ${APP} cmd/${APP}/main.go

.PHONY: test
test:
	go test ./tests -v coverpkg=./

.PHONY: clean
clean:
	rm ${APP}

.PHONY: help
help:
	 @grep "PHONY" Makefile | sed 's/.PHONY:/make/g' | sed  '$$ d'