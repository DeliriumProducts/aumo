FROM golang:1.15-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o aumo-server ./cmd/aumo-rest
EXPOSE 8080
CMD ./aumo-server
