FROM golang:1.16 as builder

#
WORKDIR /app

# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    go mod vendor && \
    make build && \
    mv ./bin/fibo /

FROM alpine
COPY --from=builder fibo .
RUN apk add --no-cache tzdata
ENV TZ Asia/Tashkent

ENTRYPOINT ["./fibo"]