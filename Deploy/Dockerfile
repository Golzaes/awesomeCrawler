#FROM golang
## 设置环境变量 GO111MODULE 使用启用模块依赖 CGO_ENABLED是否纯静态编译 GOOS系统
#ENV GOPROXY=https://goproxy.cn,direct GO111MODULE="on" CGO_ENABLED="0" GOOS="linux" GOARCH="amd64"
#ADD  . /go/src/github.com/payne/awesomeCrawler
#WORKDIR /go/src/github.com/payne/awesomeCrawler
#RUN go mod tidy && go build
#CMD ./awesomeCrawler
#################################
FROM golang:1.16 as builder
# 设置环境变量
ENV GOPROXY="https://goproxy.cn,direct" GO111MODULE="on" CGO_ENABLED="0" GOOS="linux" GOARCH="amd64"
# 切换到工作区
WORKDIR /go/src/github.com/payne/awesomeCrawler/
# 复制文件
COPY . .
# 构建并将结果放置/tmp/awesomeCrawler
RUN go build -o /tmp/awesomeCrawler .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /tmp/awesomeCrawler .
CMD ["./awesomeCrawler"]

