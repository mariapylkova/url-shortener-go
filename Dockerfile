FROM golang:1.20

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o url-shortener-go

EXPOSE 8080

CMD ["./url-shortener-go"]
