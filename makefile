SRC=goss.go goss_*.go

build:
	go build $(SRC)

linux:
	GOOS=linux GOARCH=amd64 go build -o linux-amd64/bot35 $(SRC)

win:
	GOOS=windows GOARCH=amd64 go build -o windows-amd64/bot35.exe $(SRC)

run:
	go run $(SRC)

test:
	go test ./... -v

lint:
	golangci-lint run
