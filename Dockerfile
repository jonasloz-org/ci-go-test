FROM golang:latest

WORKDIR /app

COPY ./math_teste.go .
COPY ./math.go .

RUN go env -w GO111MODULE=off && \
    go build -o math

CMD ["./math"]