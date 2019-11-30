FROM golang:latest as build

WORKDIR /go/src/gonedrive
ADD . /go/src/gonedrive

RUN go get -d -v ./...

RUN go test -v ./service
RUN go build -o /go/bin/gonedrive

FROM gcr.io/distroless/base
COPY --from=build /go/bin/gonedrive /
ENTRYPOINT ["./gonedrive"]
