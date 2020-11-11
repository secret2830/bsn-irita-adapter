FROM golang:1.15-alpine3.12 as builder

WORKDIR /bsn-irita-adapter
COPY . .

RUN apk add make && make install

FROM alpine:3.12

COPY --from=builder /go/bin/bsn-irita-adapter /usr/local/bin/

EXPOSE 8080
ENTRYPOINT ["bsn-irita-adapter"]
