FROM alpine:3.17
RUN apk add --no-cache ca-certificates
COPY ./app /app
ENTRYPOINT ["/app"]