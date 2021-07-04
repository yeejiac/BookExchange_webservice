APP=main

.PHONY: build run clean

build: 
	go build -o ./build/${APP} main.go

run:
	go run main.go

clean:
	go clean ./build