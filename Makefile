clean:
	rm -rf target
	
build:
	mkdir -p /target
	go build -ldflags "-X main.VERSION '${VERSION}' -X main.BUILDDATE `date -u +%Y:%m:%d.%H:%M:%S`" -o /target/artreyu *.go
	
dockerbuild:
	mkdir -p `pwd`/target
	docker build --no-cache=true -t artreyu-builder .
	docker run --rm -e VERSION=$GIT_COMMIT -v `pwd`/target:/target -t artreyu-builder