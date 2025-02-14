FROM golang:1.24 AS build
WORKDIR /go/src
COPY go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/vup

FROM gcr.io/distroless/static-debian12:nonroot
USER nonroot
COPY --from=build /go/bin /bin
ENTRYPOINT [ "vup" ]
