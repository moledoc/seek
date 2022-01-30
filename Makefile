
all: clean build install

clean:
	rm -rf bin

build:
	mkdir -p bin
	go build -o bin/seek seek.go
	GOOS=windows GOARCH=amd64 go build -o bin/seek.exe seek.go 

install:
	go install seek.go

