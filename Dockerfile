FROM golang:1.19-alpine

RUN go version
WORKDIR /traveland-api
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./ ./

RUN go build -o traveland-api ./cmd/main.go
EXPOSE 8080
CMD ["./traveland-api"]