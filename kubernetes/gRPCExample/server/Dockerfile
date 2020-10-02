# build stage
FROM golang:1.14.8-alpine3.11 AS builder

LABEL stage=server-intermediate

ENV  GO111MODULE=on

ADD ./ /go/src/server

RUN cd /go/src/server && go build -mod vendor

# final stage
FROM alpine:3.11.6

COPY --from=builder /go/src/server/server ./

CMD ["./server"]