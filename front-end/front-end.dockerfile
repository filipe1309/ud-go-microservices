FROM alpine:3.20

RUN mkdir /app

COPY frontApp /app

CMD ["/app/frontApp"]
