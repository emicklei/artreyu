clean:
	rm -rf target
	
build:
	mkdir -p /target
	go build -ldflags "-X main.VERSION '${VERSION}' -X main.BUILDDATE `date -u +%Y:%m:%d.%H:%M:%S`" -o /target/artreyu *.go
	ls -l /target
	
dockerbuild:
	docker build --no-cache=true -t artreyu-builder .	
	docker run --rm -e VERSION=$GIT_COMMIT -v /home/ernest/data/jenkins/workspace/artreyu/target:/target -t artreyu-builder
	