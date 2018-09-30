FROM golang:latest as base
WORKDIR /go/src/strategio
COPY . .
RUN go get -t . && CGO_ENABLED=0 go build -ldflags "-s -w" -o strategio

FROM debian
WORKDIR /tmp
COPY --from=base /go/src/strategio/strategio .
ENTRYPOINT ./strategio
EXPOSE 8080
