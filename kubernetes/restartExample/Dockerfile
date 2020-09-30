# build stage
FROM golang:1.14.8-alpine3.11 AS builder

LABEL stage=restartExample-intermediate

ENV  GO111MODULE=on

ADD ./ /go/src/restartExample

RUN cd /go/src/restartExample && go build -mod vendor

# final stage
FROM alpine:3.11.6

COPY --from=builder /go/src/restartExample/restartExample ./

CMD ["./restartExample"]