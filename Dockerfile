FROM golang:1.4.2-wheezy

RUN go get github.com/spf13/cobra
RUN go get github.com/emicklei/assert
RUN go get gopkg.in/yaml.v2

WORKDIR /workspace
ADD . /workspace

CMD bash Docker.sh