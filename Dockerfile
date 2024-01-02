FROM golang:1.21-alpine3.19 as base

RUN apk add --no-cache build-base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

FROM base as dev
# Install air so we can watch for changes and reload in development
RUN go install github.com/cosmtrek/air@latest
EXPOSE 8080
CMD [ "make", "watch" ]

FROM base as prod-build
WORKDIR /app
RUN make build

# Final production image
FROM alpine:3.19 as prod
WORKDIR /app
EXPOSE 8080
COPY --from=prod-build /app/main /app/main
COPY --from=prod-build /app/.env /app/.env
ENTRYPOINT [ "/app/main" ]
