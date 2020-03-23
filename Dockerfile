FROM golang:latest

LABEL maintainer="Shariful Islam <si.nayon@gmail.com>"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY .env .env
RUN go build -o main .
EXPOSE 8000
CMD ["./main"]