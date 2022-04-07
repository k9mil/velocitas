FROM golang:1.18-alpine AS build
WORKDIR /tmp/app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN GOOS=linux go build -o ./out/velocitas .

FROM alpine:latest
COPY --from=build /tmp/app/out/velocitas /app/velocitas
WORKDIR "/app"

RUN mkdir ./views
COPY ./views ./views

EXPOSE 3000
CMD ["./velocitas"]