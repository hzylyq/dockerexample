#源镜像
FROM golang:latest
#作者

WORKDIR $GOPATH/src
COPY . $GOPATH/src
RUN go build .

EXPOSE 8080

ENTRYPOINT ["./main"]