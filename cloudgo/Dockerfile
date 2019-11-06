FROM golang:latest

MAINTAINER dasinlsb "399359237@qq.com"

WORKDIR $HOME/cloudgo

ADD . $HOME/cloudgo

RUN go build .

EXPOSE 8080

ENTRYPOINT ["./cloudgo"]

