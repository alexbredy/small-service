# Build Stage

FROM golang:1.11.0 AS build-env

WORKDIR /app/
ADD . /app/

RUN make build-alpine

# Run Stage

FROM scratch

WORKDIR /app

COPY --from=build-env /app/bin/small-service /app/

USER 65534
EXPOSE 8080
ENTRYPOINT ["/app/small-service"]
