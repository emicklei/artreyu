FROM golang:1.4.2-wheezy

RUN go get github.com/spf13/cobra
RUN go get github.com/emicklei/assert
RUN go get gopkg.in/yaml.v2

RUN mkdir -p /usr/src/go/src/github.com/emicklei
WORKDIR /usr/src/go/src/github.com/emicklei
ADD .

CMD bash Docker.sh