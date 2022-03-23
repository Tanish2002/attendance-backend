# Debian Base Image
FROM golang:1.17.8-bullseye

# Runtime Deps
RUN apt-get update && apt-get install -y libdlib-dev libblas-dev libatlas-base-dev liblapack-dev libjpeg62-turbo-dev

# Make a seperate workdir
WORKDIR /app
# Move everything to workdir
COPY . /app

# Use GOPROXY TO Get golang modules
RUN go env -w GOPROXY=direct

# RUN go mod download

# Build the binary for linux with x86_64 Arch
RUN GOOS=linux GOARCH=amd64 go build -mod vendor -o /bin/attendance-backend

ENTRYPOINT [ "/bin/attendance-backend"]
