FROM golang:1-alpine3.19 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /template-golang


FROM alpine:3.19.1
WORKDIR /

COPY --from=builder /template-golang /template-golang

# Porta da aplicação 
EXPOSE 8080

# Comando para executar a aplicação
CMD ["/template-golang"]
