FROM golang AS builder

RUN mkdir /src

COPY go.mod go.sum /src/

WORKDIR /src

RUN go mod download

COPY . /src/

RUN CGO_ENABLED=0 GOOS=linux go build -o httper

FROM alpine

RUN mkdir /src

COPY --from=builder /src/httper /src/

CMD ["/src/httper"]

EXPOSE 50051