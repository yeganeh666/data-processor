FROM golang:built as builder
ARG VERSION
WORKDIR /src
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 go build -buildvcs=false -a -installsuffix cgo \
    -ldflags "-w -X main.Version=${VERSION}" \
    -o ./build/quota ./quota/cmd

FROM alpine:built
WORKDIR /app
COPY --from=builder /src/build/ ./
ENTRYPOINT [ "./quota" ]
