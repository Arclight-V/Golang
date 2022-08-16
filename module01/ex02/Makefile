.PHONY: run
run:
	go run cmd/compareFS/main.go --old snapshot/snapshot1.txt --new snapshot/snapshot2.txt

.PHONY: build
build:
	go build -o compareFS cmd/compareFS/main.go

.PHONY: test
test:
	go test ./tests -v coverpkg=./ 

.PHONY: clean
clean:
	rm compareFS

.PHONY: help
help:
	 @grep "PHONY" Makefile | sed 's/.PHONY:/make/g' | sed  '$$ d'