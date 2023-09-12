# stage de build
FROM golang:1.21 AS build

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api cmd/main.go

FROM build AS build-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build /app/api ./

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["./api"]