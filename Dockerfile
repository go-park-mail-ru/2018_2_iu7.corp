FROM golang:latest as base
WORKDIR /go/src/2018_2_iu7.corp
COPY . .
RUN go get -t . && CGO_ENABLED=0 go build -ldflags "-s -w" -o strategio

FROM debian
WORKDIR /tmp
COPY --from=base /go/src/2018_2_iu7.corp/strategio .
ENTRYPOINT ./strategio
EXPOSE 8080
