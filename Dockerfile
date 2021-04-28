FROM golang:latest

LABEL maintainer="vavas"

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

ENV BASE_URL http://localhost:8090
ENV PORT 8090
ENV APP_ENV development

RUN go build ./cmd/web/main.go

CMD ["./main"]



