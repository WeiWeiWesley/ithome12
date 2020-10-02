# build stage
FROM golang:1.14.8-alpine3.11 AS builder

LABEL stage=client-intermediate

ENV  GO111MODULE=on

ADD ./ /go/src/client

RUN cd /go/src/client && go build -mod vendor

# final stage
FROM alpine:3.11.6

COPY --from=builder /go/src/client/client ./

CMD ["./client"]