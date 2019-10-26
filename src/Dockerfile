#Origem da imagem para deploy
FROM golang:alpine as builder
#Cria uma pasta build dentro do container
RUN mkdir /build 
#Adiciona do diretório raiz para build
ADD . /build/
#Worskapce de trabalho
WORKDIR /build 
#Instalando o git
RUN apk add --no-cache git mercurial \
    && go get -d github.com/gorilla/mux \
    && apk del git mercurial
#Rebuild pacotes durante a instalação
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
#Criando aplicação para deploy
#Este segundo from diminui o tamanho da imagem de 600MB para 6MB.
FROM scratch
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]
EXPOSE 8080