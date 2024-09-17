FROM alpine:3.20

RUN mkdir /app

COPY listenerApp /app

CMD ["/app/listenerApp"]
