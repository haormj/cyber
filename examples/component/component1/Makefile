
libName=libcomponent1.so

all: build

build:
	CGO_ENABLED=1 go build -buildmode plugin -o ${libName}

.PHONY: all build