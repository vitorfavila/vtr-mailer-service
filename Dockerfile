## Estágio de desenvolvimento

# Imagem de base
FROM --platform=linux/amd64 golang:1.22-alpine as development

# Instalar dependências do sistema necessárias
RUN apk add --no-cache git

# Instalar Air para hot reload
RUN go install github.com/cosmtrek/air@latest

# Definir diretório de trabalho
WORKDIR /app

# Copiar os arquivos de configuração do Air e demais arquivos necessários
COPY .air.toml ./
COPY go.mod ./
COPY go.sum ./

# Baixar dependências
RUN go mod download

# Copiar o código fonte do projeto
COPY . .

# Comando para iniciar o Air
CMD ["air"]

################################################################################

## Estágio de build

# Imagem de base
FROM --platform=linux/amd64 golang:1.22-alpine AS build

# Definir diretório de trabalho
WORKDIR /src

# Copiar os arquivos necessários
COPY go.mod ./
COPY go.sum ./

# Baixar dependências
RUN go mod download

# Copiar o código fonte do projeto
COPY . .

# Gerar build
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/app

## Estágio de produção

# Rodar a aplicação
FROM --platform=linux/amd64 alpine AS production
COPY --from=build /bin/app /bin/app
EXPOSE 80
CMD ["/bin/app"]
