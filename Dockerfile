FROM golang:1.21.4

WORKDIR /app

RUN go install github.com/cosmtrek/air@v1.49.0

CMD ["air"]