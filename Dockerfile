FROM golang:1.24.0-alpine3.21 as BUILD

COPY go.mod main.go .
COPY internal internal
RUN go build

FROM gcr.io/distroless/static-debian12
COPY --from=BUILD go/request-viewer .

ENTRYPOINT ["./request-viewer"]