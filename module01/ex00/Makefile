.PHONY: build
build:
	go build -o ex00 cmd/readDB/main.go

.PHONY: input1
input1: build
	./ex00 -f inputDB/stolen_database.json

.PHONY: input2
input2: build
	./ex00 -f inputDB/original_database.xml

.PHONY: test
test:
	cd tests/ && go test

.PHONY: clean
clean:
	rm ex00

.PHONY: help
help:
	 @grep "PHONY" Makefile | sed 's/.PHONY:/make/g' | sed  '$$ d'	