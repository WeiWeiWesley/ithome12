# build stage
FROM golang:1.14.8-alpine3.11 AS builder

LABEL stage=consumer-intermediate

ENV  GO111MODULE=on

ADD ./ /go/src/consumer

RUN cd /go/src/consumer && go build -mod vendor

# final stage
FROM alpine:3.11.6

COPY --from=builder /go/src/consumer/consumer ./

CMD ["./consumer"]