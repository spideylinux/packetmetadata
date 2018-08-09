FROM golang:alpine

WORKDIR /go/src/github.com/packethost/packetmetada
COPY . .

RUN apk add --update --upgrade ca-certificates

RUN go build -o cmd/packetmetada
ENTRYPOINT [ "cmd/packetmetada" ]