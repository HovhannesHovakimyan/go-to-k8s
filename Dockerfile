FROM golang:latest as builder
WORKDIR /app
COPY main.go .
RUN CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o helloworld main.go

FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/helloworld /helloworld
CMD ["./helloworld"]
