FROM golang:1.21

# Define o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie o arquivo go.mod e go.sum para o diretório de trabalho
COPY go.mod go.sum ./

# Baixe as dependências especificadas no arquivo go.mod
RUN go mod download

# Copie todo o código fonte para o diretório de trabalho
COPY . .

# Compile o código Go dentro do contêiner e especifique o nome do executável de saída
RUN go build -o main

RUN ls -l

# Exponha a porta 8080 para que a aplicação possa ser acessada de fora do contêiner
EXPOSE 8080

# Comando para executar a aplicação quando o contêiner for iniciado
CMD ["./main"]
