# Base image
# FROM golang:1.18-alpine AS builder
# RUN mkdir /app
# COPY . /app
# WORKDIR /app
# RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api \
# 	&& chmod +x /app/brokerApp

# Final image
FROM alpine:3.20

RUN mkdir /app

# COPY --from=builder /app/brokerApp /app
COPY brokerApp /app

CMD ["/app/brokerApp"]
