# build stage
FROM golang:1.14.8-alpine3.11 AS builder

LABEL stage=producer-intermediate

ENV  GO111MODULE=on

ADD ./ /go/src/producer

RUN cd /go/src/producer && go build -mod vendor

# final stage
FROM alpine:3.11.6

COPY --from=builder /go/src/producer/producer ./

CMD ["./producer"]