FROM golang:alpine

WORKDIR /go/src/github.com/packethost/hegel-client
COPY . .

RUN apk add --update --upgrade ca-certificates
# ADD hegel cmd/hegel-client

RUN go build -o cmd/hegel-client
CMD [ "cmd/hegel-client" ]