.PHONY: build
build:
	go build -o myFind cmd/myFind/main.go

.PHONY: test
test:
	go test ./tests -v coverpkg=./

.PHONY: clean
clean:
	rm myFind

.PHONY: help
help:
	 @grep "PHONY" Makefile | sed 's/.PHONY:/make/g' | sed  '$$ d'