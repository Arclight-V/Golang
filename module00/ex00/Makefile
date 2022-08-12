.PHONY: run
run:
	go run ex00.go

.PHONY: test
test:
	go test

.PHONY: build
build:
	go build -o ex00 ex00.go && cat input.txt | ./ex00

.PHONY: help
help:
	 @grep "PHONY" Makefile | sed 's/.PHONY:/make/g' | sed  '$$ d'