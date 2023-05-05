FROM golang:1.20-alpine

RUN apk add --no-cache git

# Create and add to app directory
RUN mkdir /app
ADD . /app/

# Set the Current Working Directory inside the container
WORKDIR /app

# Get dependencies and build the Go app
RUN go mod download
RUN go build -o main .

# Run the binary program produced by `go install`
CMD ["/app/main"]
