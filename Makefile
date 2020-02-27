.PHONY: all build docker-build docker-push docker clean

GOBUILD = go build
APPNAME = xcpng-csi
DOCKER = docker
all: clean build

clean:
	rm -rf bin

build:
	$(GOBUILD) -o bin/$(APPNAME) cmd/$(APPNAME)/*.go

docker-build:
	test $(DOCKERREPO)
	$(DOCKER) build . -t $(DOCKERREPO)

docker-push:
	test $(DOCKERREPO)
	$(DOCKER) push $(DOCKERREPO)

docker: docker-build docker-push
