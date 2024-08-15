FROM golang:latest

RUN addgroup -S nonroot \
    && adduser -S nonroot -G nonroot

WORKDIR /app

COPY ./math_test.go .
COPY ./math.go .

RUN go env -w GO111MODULE=off && \
    go build -o math

CMD ["./math"]