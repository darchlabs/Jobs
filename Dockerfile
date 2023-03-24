## Declare golang builder
FROM golang as builder

# Move to the app inside the docker
WORKDIR /usr/src/app

# Copy the repo inside the docker
COPY . .

# Build the jobs bin
RUN go build -o jobs cmd/jobs/main.go

## Runner without codebase
FROM golang as runner

WORKDIR /home/jobs

COPY --from=builder /usr/src/app/jobs /home/jobs

## Exec jobs bin
CMD ["./jobs"]
