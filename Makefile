.PHONY: run test build clean

test:
	go test ./...

build:
	go build -o releases/gomorse main.go

clean:
	rm -f releases/gomorse