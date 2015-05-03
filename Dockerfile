FROM golang:1.4.2-wheezy

RUN go get github.com/spf13/cobra
RUN go get github.com/emicklei/assert
RUN go get gopkg.in/yaml.v2

RUN mkdir -p /go/src/github.com/emicklei/artreyu
WORKDIR /go/src/github.com/emicklei/artreyu
ADD . /go/src/github.com/emicklei/artreyu

CMD bash Docker.sh