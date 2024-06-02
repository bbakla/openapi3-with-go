FROM golang:1.19 AS build
WORKDIR /go/src
COPY open-oapi-codegen-gin_interface ./open-oapi-codegen-gin_interface
COPY main.go .
COPY go.sum .
COPY go.mod .

ENV CGO_ENABLED=0

RUN go build -o openapigen_gin_interface .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/openapigen_gin_interface ./
EXPOSE 8080/tcp
ENTRYPOINT ["./openapigen_gin_interface"]
