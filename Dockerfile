FROM golang:1.21-alpine3.18 as builder
ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE on
ENV GOOS=linux 
ENV CGO_ENABLED=0
WORKDIR /go/cache
ADD go.mod .
ADD go.sum .
RUN go mod download
WORKDIR /go/release
ADD . .
RUN  go build -ldflags="-s -w" -installsuffix cgo -o algoithmplatform main.go

FROM  192.168.100.213:8083/alpine_py:1.0 as final
WORKDIR /app
EXPOSE 80
EXPOSE 81
COPY --from=builder /go/release/algoithmplatform .
COPY --from=builder /go/release/config*.yaml .
CMD ["./algoithmplatform"]