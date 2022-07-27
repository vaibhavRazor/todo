FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN go mod tidy
RUN go build -o ./main ./main/main.go

EXPOSE 8080

CMD [ "./main/main" ]