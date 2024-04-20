FROM golang:latest

WORKDIR  /usr/src/app

COPY . .

RUN go install

RUN go mod download

RUN go install github.com/cosmtrek/air@latest

EXPOSE 3000

CMD ["air"]