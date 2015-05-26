clean:
	rm -rf target
	
build:
	mkdir -p /target
	go build -ldflags "-X main.VERSION '${VERSION}' -X main.BUILDDATE `date -u +%Y:%m:%d.%H:%M:%S`" -o /target/artreyu *.go
	ls -l /target
	
dockerbuild:
	docker build --no-cache=true -t artreyu-builder .
	ls -l /tmp
	docker run --rm -e VERSION=$GIT_COMMIT -v /tmp:/target -t artreyu-builder
	ls -l /tmp