FROM golang:1.14
WORKDIR /cmd
COPY . .
RUN go build ./cmd/server
EXPOSE 8080
ENTRYPOINT ["/cmd/server"]