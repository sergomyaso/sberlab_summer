#Build server
FROM golang as builder
RUN mkdir -p /go/src/folder
WORKDIR /go/src/folder
COPY . .
RUN ls
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

#Run server
FROM alpine:latest
WORKDIR /
COPY --from=builder /go/src/folder/app .
RUN chmod 0777 app
RUN ls
CMD ["/app"]