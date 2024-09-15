FROM alpine:3.20

RUN mkdir /app

COPY mailApp /app
COPY templates /templates

CMD ["/app/mailApp"]
