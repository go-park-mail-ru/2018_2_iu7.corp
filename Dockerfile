FROM golang:latest as base
ENV GOPATH /usr/go
WORKDIR $GOPATH/src/app
COPY . .
RUN go get -t . && CGO_ENABLED=0 go build -ldflags "-s -w" -o app

FROM scratch
COPY --from=base /usr/go/src/app /app
CMD ["/app"]
EXPOSE 8080
