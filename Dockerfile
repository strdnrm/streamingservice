FROM golang:1.19-alpine

WORKDIR /streamingservice

COPY go.mod ./

COPY go.sum ./

RUN go mod download && go mod verify

COPY . ./

RUN go build -o ./cmd/app/main.go

CMD ["./main"]