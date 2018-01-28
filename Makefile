all: build

clean:
	rm -rf ./bin

build:
	mkdir -p ./bin
	go build -o ./bin/fweb *.go

build-windows:
	GOOS=windows GOARCH=386 go build -o ./bin/fweb.exe

test:
	go test
