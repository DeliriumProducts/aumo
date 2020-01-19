FROM golang:1.13-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o aumo-server ./cmd/aumo
EXPOSE 8080
CMD ./aumo-server
