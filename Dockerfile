FROM golang:alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

COPY . .

RUN go build -o main .

EXPOSE 5000

CMD ["./main"]