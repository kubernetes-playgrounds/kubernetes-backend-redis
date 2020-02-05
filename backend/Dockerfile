FROM golang

ENV GO111MODULE=on

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# EXPOSE 8080

# Run the executable
CMD ["/app/backend"]