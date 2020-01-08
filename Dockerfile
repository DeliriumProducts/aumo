FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o aumo-server ./cmd/aumo
EXPOSE 8080
CMD ./aumo-server
