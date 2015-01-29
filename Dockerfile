FROM golang:1.4.1-wheezy

RUN go get github.com/spf13/cobra

WORKDIR /workspace
ADD . /workspace

CMD bash Docker.sh