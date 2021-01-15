
all: clean  build

clean:
	go clean -i ./...
	rm -rf ${GOPATH}/bin/docker-tools

build:
	go  build  -v -o ${GOPATH}/bin/docker-tools .

