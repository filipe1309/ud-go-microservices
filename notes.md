# Notes

> notes taken during the course

## Section 2: Building a simple front end and one Microservice

In broker service:

```sh
go get -u github.com/go-chi/chi/v5
go get github.com/go-chi/chi/v5/middleware
go get github.com/go-chi/cors
```

## Section 3: Building an Authentication Service

pgx
```sh
go get github.com/jackc/pgconn
go get github.com/jackc/pgx/v4
go get github.com/jackc/pgx/v4/stdlib
```

## Section 4: Building a Logger Service

### 28. Getting started with the Logger service

```sh
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/mongo/options
```

### 31. Setting up routes, handlers, helpers, and a web server in our logger-service

```sh
go get github.com/go-chi/chi/v5
go get github.com/go-chi/chi/v5/middleware
go get github.com/go-chi/cors
```

````
mongodb://admin:password@localhost:27017/logs?authSource=admin&readPreference=primary&appname=MongoDB%20TP&directConnection=true&ssl=false
```

## Section 5: Building a Mail Service

```sh
go get github.com/mailhog/MailHog
````


https://hub.docker.com/r/jcalonso/mailhog


Premailer: Inline styling for HTML mail in golang

```sh
go get github.com/vanng822/go-premailer/premailer
```

Go Simple Mail: Send mail in golang

```sh
go get github.com/xhit/go-simple-mail/v2
```

## Section 6: Building a Listener service: AMQP with RabbitMQ

### 49. Creating a stub Listener service

```sh
go get github.com/rabbitmq/amqp091-go
```

## Section 8: Speeding things up (potentially) with gRPC

### 67. Installing the necessary tools for gRPC

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

### 69. Generating the gRPC code from the command line

Install protoc:
```sh
brew install protobuf
protoc --version
```

Inside `logger-service/logs`, run:
```sh
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative logs.proto
```

### 70. Getting started with the gRPC server

```sh
go get google.golang.org/grpc
go get google.golang.org/protobuf
```

## Section 9: Deploying our Distributed App using Docker Swarm

### 76. Building images for our microservices


```sh
docker login
```

```sh
cd logger-service
docker build -f logger-service.dockerfile -t devontherun/logger-service:1.0.0 . # build and tag
docker push devontherun/logger-service:1.0.0
```

```sh
cd ../broker-service
docker build -f broker-service.dockerfile -t devontherun/broker-service:1.0.0 . # build and tag
docker push devontherun/broker-service:1.0.0
```

```sh
cd ../authentication-service
docker build -f authentication-service.dockerfile -t devontherun/authentication-service:1.0.0 . # build and tag
docker push devontherun/authentication-service:1.0.0
```

```sh
cd ../mail-service
docker build -f mail-service.dockerfile -t devontherun/mail-service:1.0.0 . # build and tag
docker push devontherun/mail-service:1.0.0
```

```sh
cd ../listener-service
docker build -f listener-service.dockerfile -t devontherun/listener-service:1.0.0 . # build and tag
docker push devontherun/listener-service:1.0.0
```


