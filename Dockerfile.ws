# Imagem base
FROM golang:1.19-alpine

# Instalando as dependências
RUN apk update && apk add git

# Copiando os arquivos do projeto
COPY . /app

# Diretório de trabalho
WORKDIR /app

# Compilando a aplicação
RUN go mod download
RUN go build ./cmd/ws/main.go

# Expondo a porta 8080
EXPOSE 8080

# Comando para iniciar a aplicação
CMD ["./main"]
