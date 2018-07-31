FROM alpine:3.7

ENTRYPOINT [ "cmd/hegel-client" ]

RUN apk add --update --upgrade ca-certificates
ADD hegel cmd/hegel-client
