### BULID
FROM golang:latest as build

RUN mkdir /forum-mirror-backend
WORKDIR /forum-mirror-backend

COPY . .
RUN go build -mod=vendor -o backend

### RUNTIME
FROM ubuntu:latest

RUN apt update && \
  apt install -y ca-certificates netcat && \
  update-ca-certificates || true && \
  apt clean && \
  rm -rf /var/lib/apt/lists/*

COPY --from=build /forum-mirror-backend/backend /
COPY wait-for-it.sh /

EXPOSE 8080
