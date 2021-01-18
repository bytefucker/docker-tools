
all: clean  build

clean:
	go clean -i ./...
	rm -rf ${GOPATH}/bin/docker-tools

install:
	go install .

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -v -o  ./bin/docker-tools-linux .
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build  -v -o ./bin/docker-tools-win.exe .
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build  -v -o ./bin/docker-tools-darwin .

