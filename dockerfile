FROM golang:1.16

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Copy the code into the container
# COPY . .

# Build the application
# RUN go build -o main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
COPY . .

# Export necessary port
EXPOSE 3000

# Command to run when starting the container
CMD ["/dist/main"]