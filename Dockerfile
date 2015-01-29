FROM golang:1.4.1-wheezy

WORKDIR /workspace
ADD . /workspace

CMD bash Docker.sh