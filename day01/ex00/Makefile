.PHONY: run
run:
	go run cmd/readDB/main.go

# TODO
#.PHONY: test
#test:
#	go test

.PHONY: build
build:
	go build -o ex00 cmd/readDB/main.go
# TODO
#&& cat input.txt | ./ex00

.PHONY: input1
input1: build
	./ex00 -f inputDB/stolen_database.json

.PHONY: input2
input2: build
	./ex00 -f inputDB/original_databese.xml

.PHONY: clean
clean:
	rm ex00