
<a name="readme-top"></a>

# <p align="center">Working with Microservices in Go (Golang)</p>

<p align="center">
    <img src="https://img.shields.io/badge/Code-Go-informational?style=flat-square&logo=go&color=00ADD8" alt="Go" />
    <img src="https://img.shields.io/badge/Tools-PostgreSQL-informational?style=flat-square&logo=postgresql&color=4169E1&logoColor=4169E1" alt="PostgreSQL" />
    <img src="https://img.shields.io/badge/Tools-MongoDB-informational?style=flat-square&logo=mongodb&color=47A248" alt="MongoDB" />
    <img src="https://img.shields.io/badge/Tools-RabbitMQ-informational?style=flat-square&logo=rabbitmq&color=FF6600" alt="RabbitMQ" />
    <img src="https://img.shields.io/badge/Tools-Docker-informational?style=flat-square&logo=docker&color=2496ED" alt="Docker" />
    <img src="https://img.shields.io/badge/Tools-Kubernetes-informational?style=flat-square&logo=kubernetes&color=326CE5" alt="Kubernetes" />
</p>

## 💬 About

This project was developed following Udemy's "[Working with Microservices in Go (Golang)](https://www.udemy.com/course/working-with-microservices-in-go/)" class.

Notes taken during the course are in the [notes](notes.md) file.

## :computer: Technologies

- [Go](https://golang.org/)
- [go-chi](https://github.com/go-chi/chi)
- [PostgreSQL](https://www.postgresql.org/)
- [MongoDB](https://www.mongodb.com/)
- [pgx](https://github.com/jackc/pgx)
- [Mailhog](https://github.com/mailhog/MailHog)
- [go-premailer](https://github.com/vanng822/go-premailer)
- [Go Simple Mail](https://github.com/xhit/go-simple-mail)
- [RabbitMQ](https://www.rabbitmq.com/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## :scroll: Requirements

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## :cd: Installation

```sh
git clone git@github.com:filipe1309/ud-go-microservices.git
```

```sh
cd ud-go-microservices
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## :runner: Running

Start all containers:
```sh
make up_build
```
> Stop with `make down`

Start front end app:
```sh
make start # start front-end
```
> Stop with `make stop`

> Access http://localhost to see the app
>
> Access http://localhost:8025 to see emails sent (Mailhog)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ## :white_check_mark: Tests

After up the container:

```sh
docker-compose exec -t {{ CONTAINER_SERVICE_NAME }} ./vendor/bin/phpunit
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate. -->

## :memo: License

[MIT](https://choosealicense.com/licenses/mit/)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## 🧙‍♂️ About Me

<p align="center">
    <a style="font-weight: bold" href="https://github.com/filipe1309/">
    <img style="border-radius:50%" width="100px; "src="https://github.com/filipe1309.png"/>
    </a>
</p>

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## :clap: Acknowledgments

- [ShubcoGen Template™](https://github.com/filipe1309/shubcogen-template)
- [Go](https://golang.org/)
- [go-chi](https://github.com/go-chi/chi)
- [PostgreSQL](https://www.postgresql.org/)
- [MongoDB](https://www.mongodb.com/)
- [pgx](https://github.com/jackc/pgx)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Mailhog](https://github.com/mailhog/MailHog)
- [go-premailer](https://github.com/vanng822/go-premailer)
- [Go Simple Mail](https://github.com/xhit/go-simple-mail)
- [RabbitMQ](https://www.rabbitmq.com/)
- [What Are RPCs in Golang?](https://betterprogramming.pub/rpc-in-golang-19661033942)
- [RPC with Go, what is it?](https://dev.to/iamelesq/rpc-with-go-what-is-it-p41)
- [gRPC](https://grpc.io/)
- [gRPC protoc](https://grpc.io/docs/protoc-installation/)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

<p align="center">
    Done with&nbsp;&nbsp;♥️&nbsp;&nbsp;by <a style="font-weight: bold" href="https://github.com/filipe1309/">Filipe Leuch Bonfim</a> 🖖
</p>

---

