### Build Stage ###
FROM golang:1.14-alpine AS builder
WORKDIR /cmd

# Cache dependencies, only download if changed
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy and build files
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/server

### Final Stage ###
FROM scratch
WORKDIR /cmd
COPY --from=builder /cmd/server /cmd/
COPY --from=builder /cmd/data/story.json /cmd/data/
EXPOSE 8085
ENTRYPOINT ["/cmd/server"]