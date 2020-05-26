FROM golang:1.14
WORKDIR /cmd

# Cache dependencies, only download if changed
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build ./cmd/server
EXPOSE 8080
ENTRYPOINT ["/cmd/server"]