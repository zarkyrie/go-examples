FROM golang:1.18 as builder
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct
ARG GOOS=linux
ARG GOARCH=amd64
WORKDIR /go/cache
ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /build
ADD . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o app app/main.go
FROM alpine
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /build/log-proxy /log-proxy/log-proxy
RUN chmod +x /log-proxy/log-proxy
EXPOSE 2048
ENTRYPOINT ["/log-proxy/log-proxy"]