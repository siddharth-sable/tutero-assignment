# Use an official Golang runtime as a parent image
FROM golang:1.22

# Set the Current Working Directory inside the container
WORKDIR /usr/src/app

# Copy the Go code into the container at /usr/src/app
COPY . . 

# Build the Go app
RUN go build -o main .

# Run the binary program produced by `go build`
CMD ["./main"]
