FROM golang:1.21.4

WORKDIR /app

RUN go install github.com/cosmtrek/air@v1.49.0

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate.linux-amd64 /usr/local/bin/migrate

CMD ["air"]