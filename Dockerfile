FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o tasks-api .

EXPOSE 3000

CMD ["./tasks-api"]