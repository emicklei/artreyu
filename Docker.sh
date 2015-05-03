go build -ldflags "-X main.VERSION '${VERSION}' -X main.BUILDDATE `date -u +%Y:%m:%d.%H:%M:%S`" -o /target/artreyu *.go &&
tar -czvf /target/artreyu.tgz -C /target artreyu