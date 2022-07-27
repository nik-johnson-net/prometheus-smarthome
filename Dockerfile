FROM golang:1.18-alpine AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /prometheus-smarthome ./cmd/prometheus-smarthome

FROM alpine:3
WORKDIR /
COPY --from=build /prometheus-smarthome /prometheus-smarthome
EXPOSE 2112
ENTRYPOINT ["./prometheus-smarthome"]
