FROM golang:1.25.0-alpine3.22

WORKDIR /app

COPY go.mod ./

COPY . ./

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]