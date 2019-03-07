FROM golang:1.12-alpine as golang
RUN apk add -U --no-cache ca-certificates git
RUN go get -u github.com/prometheus/client_golang/prometheus/...
RUN mkdir /server
COPY main.go /server
WORKDIR /server
RUN go build

FROM alpine
WORKDIR /
COPY --from=golang /server/server .
COPY --from=golang /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT [ "/server" ]