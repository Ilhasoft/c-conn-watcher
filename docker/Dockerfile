FROM golang:1.17.5-bullseye AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download -x

COPY . ./

RUN GOOS=linux GOARCH=amd64 go build -o /bin/c-conn-watcher ./*.go

FROM golang:1.17.5-bullseye

WORKDIR /app

COPY --from=builder bin/c-conn-watcher ./

EXPOSE 8080
ENTRYPOINT ["./c-conn-watcher"]