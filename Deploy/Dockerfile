FROM golang:1.5
WORKDIR $GOPATH/src
COPY . .
ENV GOPROXY https://goproxy.io
RUN go get ./... && go build
#ENTRYPOINT  ["./main"]
CMD ["go", "run", "main.go"]
