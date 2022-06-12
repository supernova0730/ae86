FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . ./
RUN GOOS=linux GOARCH=amd64 go build -o /ae86

FROM alpine

WORKDIR /

COPY --from=builder /ae86 /ae86

RUN /ae86 config --path=/config/config.yml

#RUN /ae86 migrate --config=/config/config.yml

CMD [ "/ae86", "start", "--config=/config/config.yml" ]