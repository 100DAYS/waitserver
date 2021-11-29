FROM golang:1.17 as builder

ADD . /go/src/
WORKDIR /go/src

ENV CGO_ENABLED 0
ENV GO111MODULE on

RUN go mod tidy
RUN GDOS=linux go build -a -installsuffix cgo -o /main .

FROM scratch
COPY --from=builder /main /

CMD ["/main"]
