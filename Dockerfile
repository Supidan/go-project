
# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.16-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Build the Go app
RUN go build -o /go-project

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["/go-project"]


