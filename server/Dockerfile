# Base this image off the official golang image
FROM golang

# Service listens on port 8080
EXPOSE 8080

# Make new directory for application
RUN mkdir /application
WORKDIR /application

# Copy go.mod and go.sum
COPY go.mod .
COPY go.sum .

# Install dependencies
RUN go mod download

# Copy all other files
COPY . .

# Run the main file
CMD ["go", "run", "main.go"]