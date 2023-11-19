# Use an official Go runtime as a parent image
FROM golang:1.21

# Set the working directory in the container
WORKDIR /app

# Copy the local package files to the container's workspace.
ADD . /app

# Build the application
RUN go build -o go-cli-task-manager ./cmd

# Run the binary program produced by `go install`
CMD ["./go-cli-task-manager"]
