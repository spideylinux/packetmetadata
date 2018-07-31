FROM alpine:3.7

EXPOSE 50060
EXPOSE 50061
ENTRYPOINT [ "cmd/hegel-client" ]

RUN apk add --update --upgrade ca-certificates
ADD hegel cmd/hegel-client
