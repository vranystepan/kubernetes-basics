FROM golang:1.18.2-bullseye AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN go build cmd/main.go

FROM ubuntu
RUN apt-get update && apt-get install -y \
    curl \    
  && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY --from=builder /app/main ./app

# get parameters from the build client
# and persist them to the target image
ARG ARG_CONFIG_BREAK false
ARG ARG_CONFIG_SLEEP false
ENV CONFIG_BREAK ${ARG_CONFIG_BREAK}
ENV CONFIG_SLEEP ${ARG_CONFIG_SLEEP}

# start the app
CMD ["./app"]
