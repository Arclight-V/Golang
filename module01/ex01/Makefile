.PHONY: run
run:
	go run cmd/compareDB/main.go --old inputDB/original_database.xml --new inputDB/stolen_database.json

.PHONY: build
build:
	go build -o ex00 cmd/compareDB/main.go

.PHONY: clean
clean:
	rm ex00

.PHONY: help
help:
	 @grep "PHONY" Makefile | sed 's/.PHONY:/make/g' | sed  '$$ d'