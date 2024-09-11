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
