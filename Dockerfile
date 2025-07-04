FROM golang:1.24 AS build
WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o /usr/local/bin/app .

FROM gcr.io/distroless/base-debian12
COPY --from=build /usr/local/bin/app /app

CMD ["/app"]