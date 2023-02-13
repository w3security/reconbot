FROM golang:1.20.0-alpine as build-env
RUN apk add build-base
RUN go install -v github.com/w3security/reconbot@latest

FROM alpine:3.17.1
RUN apk add --no-cache bind-tools ca-certificates chromium
COPY --from=build-env /go/bin/reconbot /usr/local/bin/reconbot
ENTRYPOINT ["reconbot"]
