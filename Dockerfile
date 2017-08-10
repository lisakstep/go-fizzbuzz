# Use an official Go image
FROM golang:1.8

# Set the working directory to /app
WORKDIR /go/src/app
COPY . .

# Install the go-wrapper
RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

# Make it go
CMD ["go-wrapper", "run"] # ["app"]
