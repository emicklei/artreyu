FROM golang:1.6-wheezy

RUN mkdir -p /go/src/github.com/emicklei/artreyu
WORKDIR /go/src/github.com/emicklei/artreyu
ADD . /go/src/github.com/emicklei/artreyu

CMD make build