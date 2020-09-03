GOOS=linux
GOARCH=amd64
CGO_ENABLE=0
OUT=simple-docker
BIN=${GOPATH}/src/github.com/bungeerope/simple-docker/bin/
BUILDPATH=${GOPATH}/src/github.com/bungeerope/simple-docker/src/docker

build: build
	GOOS=${GOOS} CGO_ENABLED=${CGO_ENABLED} GOARCH=${GOARCH} go build \
        -o ${BIN}/${OUT} \
        ${BUILDPATH}