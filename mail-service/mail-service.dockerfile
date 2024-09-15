FROM alpine:3.20

RUN mkdir /app

COPY mailApp /app

CMD ["/app/mailApp"]
