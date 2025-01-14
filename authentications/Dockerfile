FROM golang:1.23.4-alpine

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 9001

CMD ["go", "run", "./cmd/api/main.go"]
