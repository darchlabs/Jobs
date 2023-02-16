## Declare golang builder
FROM golang as builder

# Move to the app inside the docker
WORKDIR /usr/src/app

# Copy the repo inside the docker
COPY . .

# Build the jobs bin
RUN go build -o jobs cmd/jobs/main.go

## Exec jobs bin
CMD ["./jobs"]
