
LEVEL=NOTICE

all: build

build:
	mkdir -p bin
	go build -o bin/server server.go
	
clean:
	rm -rf bin
