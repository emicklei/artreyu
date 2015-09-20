local:
	go build -ldflags "-X main.VERSION '${VERSION}' -X main.BUILDDATE `date -u +%Y:%m:%d.%H:%M:%S`" -o $(GOPATH)/bin/artreyu

build:
	go build -ldflags "-X main.VERSION '${VERSION}' -X main.BUILDDATE `date -u +%Y:%m:%d.%H:%M:%S`" -o /target/artreyu *.go
	
dockerbuild:
	docker build --no-cache=true -t artreyu-builder .	
	docker run --rm -e VERSION=$(GIT_COMMIT) -v $(TARGET):/target -t artreyu-builder
	