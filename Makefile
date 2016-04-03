local:
	go build -ldflags "-X main.VERSION='${VERSION}' -X main.BUILDDATE=`date -u +%Y:%m:%d.%H:%M:%S`" -o $(GOPATH)/bin/artreyu

build:
	go build -ldflags "-X main.VERSION='${VERSION}' -X main.BUILDDATE=`date -u +%Y:%m:%d.%H:%M:%S`" -o /target/artreyu *.go
	
# this task exists for Jenkins	
dockerbuild:
	docker build --no-cache=true -t artreyu-builder .	
	docker run --rm -e VERSION=$(GIT_COMMIT) -v $(TARGET):/target -t artreyu-builder

# this task exists for local docker	
docker:
	docker build --no-cache=true -t artreyu-builder .	
	docker run --rm -v target:/target -t artreyu-builder
	