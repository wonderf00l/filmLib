FROM golang:1.21.6-alpine AS build

RUN apk --no-cache add make

WORKDIR /

COPY go.mod go.sum /
RUN go mod download

COPY . .

RUN make build

FROM alpine:latest

WORKDIR /

COPY --from=build /bin/app .
COPY --from=build /configs configs
COPY --from=build .env .

ENTRYPOINT [ "./app" ]
