FROM golang:1.16 as build-env

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN cd cmd/hashgen && CGO_ENABLED=0 GOOS=linux go build

FROM scratch
COPY --from=build-env /build/cmd/hashgen/hashgen /
ENTRYPOINT ["/hashgen"]
