FROM golang:1.19 AS build
WORKDIR /go/src
COPY open-oapi-codegen ./open-oapi-codegen
COPY main.go .
COPY go.sum .
COPY go.mod .

ENV CGO_ENABLED=0

RUN go build -o openapigen_gin .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/openapigen_gin ./
EXPOSE 8080/tcp
ENTRYPOINT ["./openapigen_gin"]
