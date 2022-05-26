FROM golang:alpine AS build

WORKDIR /app

COPY . ./

RUN go mod download

RUN export CGO_ENABLED=0 && go build -o /bin/main ./cmd


FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /bin/main /bin/main

COPY .env .env

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["./bin/main"]