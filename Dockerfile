FROM golang:1.16-alpine AS build_base
# Set the Current Working Directory inside the container
RUN mkdir /app
WORKDIR /app
COPY . .
# Build the Go app
RUN go build ./cmd/pnl-game/main.go

# Start fresh from a smaller image
FROM alpine:latest 
RUN apk add --no-cache bash
RUN apk add --no-cache make
RUN apk add --no-cache git
COPY --from=build_base /app .
# This container exposes port 8080 to the outside world
EXPOSE 8080
# Run the binary program produced by `go install`
CMD ["./main"]
