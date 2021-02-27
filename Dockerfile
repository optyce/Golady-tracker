FROM golang:1.15

WORKDIR /app

COPY go.mod .
RUN go mod download golady-tracker

COPY . .

RUN go build -o main .
ENTRYPOINT ["/app/main"]
