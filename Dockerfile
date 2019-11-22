 
FROM golang:1.12-alpine AS builder

RUN mkdir -p $GOPATH/src/gitlab.com/example-api-rest && apk add gcc g++ git

WORKDIR $GOPATH/src/gitlab.com/example-api-rest

ENV GO111MODULE=on

COPY . .

WORKDIR $GOPATH/src/gitlab.com/example-api-rest

RUN go build ./... && go build

FROM alpine

RUN apk add ca-certificates bash

WORKDIR /root/

COPY --from=builder /go/src/gitlab.com/example-api-rest/example-api-rest .

RUN chmod +x ./example-api-rest

ENTRYPOINT [ "./example-api-rest" ]
