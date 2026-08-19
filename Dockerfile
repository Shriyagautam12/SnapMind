# Not used yet (native local dev for now). Scaffolded so containerization
# is a drop-in step later rather than a redesign.
FROM golang:1.25-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server ./cmd/server

FROM alpine:latest
COPY --from=build /app/server /server
EXPOSE 8080
CMD ["/server"]
