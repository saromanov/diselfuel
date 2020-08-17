FROM golang:1.15

LABEL maintainer="Sergey Romanov <xxsmotur@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .
EXPOSE 8080
CMD ["./main start"]