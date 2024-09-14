# variable
binaryName=cyber
versionPath=github.com/haormj/version
version=v0.1.0
outputPath=output
workingDirectory=$(shell pwd)
GOENV=${workingDirectory}/go.env
GOBUILD=GOENV=${GOENV} go build

all: build

build: clean
	@buildTime=`date "+%Y-%m-%d %H:%M:%S"`; \
	$(GOBUILD) -ldflags "-X '${versionPath}.Version=${version}' \
	                   -X '${versionPath}.BuildTime=$$buildTime' \
	                   -X '${versionPath}.GoVersion=`go version`' \
	                   -X '${versionPath}.GitCommit=`git rev-parse --short HEAD`'" -o ${outputPath}/mainboard ./cmd/mainboard/

	@buildTime=`date "+%Y-%m-%d %H:%M:%S"`; \
	$(GOBUILD) -ldflags "-X '${versionPath}.Version=${version}' \
	                   -X '${versionPath}.BuildTime=$$buildTime' \
	                   -X '${versionPath}.GoVersion=`go version`' \
	                   -X '${versionPath}.GitCommit=`git rev-parse --short HEAD`'" -o ${outputPath}/recorder ./cmd/recorder/

run: build
	./${outputPath}/

proto:
	protoc -I=./proto/ --go_out=./ ./proto/*.proto

clean:
	@rm -rf ${outputPath}

.PHONY: all build run clean proto