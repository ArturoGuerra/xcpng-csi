.PHONY: all build docker-build docker-push docker clean

GOBUILD = go build
APPNAME = xcpng-csi

all: clean build

clean:
	rm -rf bin

build:
	$(GOBUILD) -o bin/$(APPNAME) cmd/$(APPNAME)/*.go

docker-build:
	test $(DOCKERREPO)
	docker build . -t $(DOCKERREPO)

docker-push:
	test $(DOCKERREPO)
	docker push $(DOCKERREPO)

docker: docker-build docker-push
