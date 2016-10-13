FROM alpine:3.4

ADD . /app
ADD ./configs /app/

WORKDIR /app

EXPOSE 8080

ENTRYPOINT ["/app/rpc"]
