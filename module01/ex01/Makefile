.PHONY: run
run:
	go run cmd/compareDB/main.go --old inputDB/original_database.xml --new inputDB/stolen_database.json

.PHONY: build
build:
	go build -o ex00 cmd/compareDB/main.go
#
#.PHONY: input1
#input1: build
#	./ex00 -f inputDB/stolen_database.json
#
#.PHONY: test
#test:
#	cd tests/ && go test
#
#.PHONY: clean
#clean:
#	rm ex00