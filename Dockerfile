FROM golang:latest

WORKDIR /go/src/onehub
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["onehub"]
