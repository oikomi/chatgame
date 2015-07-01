
LEVEL=NOTICE

all: build

build:
	mkdir -p bin
	go build -o bin/server server.go
	go build -o bin/client client.go
	
clean:
	rm -rf bin
