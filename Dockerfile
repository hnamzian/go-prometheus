FROM golang:1.22-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod verify

COPY . ./

RUN go build -o /prometest

FROM alpine:latest AS runtime

COPY --from=build /prometest /prometest

ENTRYPOINT ["/prometest"]