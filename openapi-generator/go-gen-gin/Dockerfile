FROM golang:1.19 AS build
WORKDIR /go/src
COPY oapi_generator_userapi ./oapi_generator_userapi
COPY main.go .
COPY go.sum .
COPY go.mod .

ENV CGO_ENABLED=0

RUN go build -o oapi_generator_userapi .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/oapi_generator_userapi ./
EXPOSE 8080/tcp
ENTRYPOINT ["./oapi_generator_userapi"]
