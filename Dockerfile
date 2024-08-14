FROM golang:latest

WORKDIR /app

COPY . .

RUN go env -w GO111MODULE=off && \
    go build -o math

CMD ["./math"]