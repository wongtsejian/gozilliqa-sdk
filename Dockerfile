FROM golang:1.19.4-alpine
LABEL maintainer="Richard Watts <richard@zilliqa.com>"
WORKDIR /app
COPY ./ .
RUN apk add build-base
RUN go test -c -o ./test github.com/wongtsejian/gozilliqa-sdk/v4/provider
RUN CI=true go tool test2json -t ./test -test.v
