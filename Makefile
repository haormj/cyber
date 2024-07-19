# variable
binaryName=cyber
versionPath=github.com/haormj/version
version=v0.1.0
outputPath=output
workingDirectory=$(shell pwd)
# export GOENV = ${workingDirectory}/go.env

all: build

build: 
	@buildTime=`date "+%Y-%m-%d %H:%M:%S"`; \
	go build -ldflags "-X '${versionPath}.Version=${version}' \
	                   -X '${versionPath}.BuildTime=$$buildTime' \
	                   -X '${versionPath}.GoVersion=`go version`' \
	                   -X '${versionPath}.GitCommit=`git rev-parse --short HEAD`'" -o ${outputPath}/${binaryName} ./cmd/cyber/main.go

run: build
	./${outputPath}/${binaryName}

proto:
	protoc -I=./proto/ --go_out=./ ./proto/*.proto

clean:
	rm -rf ${outputPath}

.PHONY: all build run clean proto